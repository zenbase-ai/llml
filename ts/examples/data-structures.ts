import { llml } from "../src/index";

/**
 * Complex Data Structures Example: Hierarchical Organizational Analysis
 *
 * Demonstrates how LLML transforms deeply nested organizational data into
 * structured prompts for AI-driven insights and analysis. Shows the pattern:
 * Complex Data Fetching → LLML Transformation → AI Analysis
 */

export async function analyzeOrganizationalStructure(
	departmentId: string,
): Promise<{
	department: string;
	analysis: string;
	timestamp: string;
	dataPoints: {
		teamsAnalyzed: number;
		projectsEvaluated: number;
		totalMembers: number;
		budgetScope: number;
	};
}> {
	// 1. Fetch data from external services
	const orgChart = await orgService.getDepartment({ id: departmentId });
	const projects = await projectTracker.getActiveProjects({
		_department: orgChart.name,
	});
	const performance = await performanceTracker.getTeamMetrics({
		_department: orgChart.name,
	});

	// 2. Transform with LLML - structure complex hierarchical data
	const analysisPrompt = llml({
		task: "analyze organizational structure, resource allocation, and identify optimization opportunities",
		organization: {
			department: {
				name: orgChart.name,
				leadership: {
					head: orgChart.head.name,
					title: orgChart.head.title,
					experience: orgChart.head.experience,
				},
				structure: orgChart.teams.map((team) => ({
					name: team.name,
					lead: team.lead,
					size: team.members.length,
					composition: {
						senior: team.members.filter((m) => m.seniority === "senior").length,
						mid: team.members.filter((m) => m.seniority === "mid").length,
						junior: team.members.filter((m) => m.seniority === "junior").length,
					},
					technologies: team.technologies,
					financials: {
						budget: team.budget,
						headcountGap: team.headcount.planned - team.headcount.current,
					},
				})),
				aggregates: {
					totalMembers: orgChart.metrics.totalHeadcount,
					growthPlan:
						orgChart.metrics.plannedHeadcount - orgChart.metrics.totalHeadcount,
					budgetUtilization: orgChart.metrics.utilizationRate,
					averageTenure: orgChart.metrics.averageTenure,
				},
			},
			workload: {
				activeProjects: projects.map((project) => ({
					name: project.name,
					team: project.assignedTeam,
					status: project.status,
					priority: project.priority,
					progress: `${Math.round(project.timeline.progress * 100)}%`,
					resourceAllocation: project.assignedMembers.length,
					budgetHealth: {
						utilization: project.budget.spent / project.budget.allocated,
						remaining: project.budget.remaining,
						onTrack:
							project.budget.spent / project.budget.allocated <=
							project.timeline.progress,
					},
					riskProfile: {
						riskCount: project.risks.length,
						highRisks: project.risks.filter((r) => r.severity === "high")
							.length,
						blockers: project.status === "blocked" ? project.dependencies : [],
					},
				})),
				projectDistribution: {
					byPriority: {
						high: projects.filter((p) => p.priority === "high").length,
						medium: projects.filter((p) => p.priority === "medium").length,
						low: projects.filter((p) => p.priority === "low").length,
					},
					byStatus: {
						active: projects.filter((p) => p.status === "in-progress").length,
						planning: projects.filter((p) => p.status === "planning").length,
						blocked: projects.filter((p) => p.status === "blocked").length,
					},
				},
			},
			performance: {
				teamVelocity: Object.entries(performance.productivity.velocity).map(
					([team, metrics]) => ({
						team,
						current: metrics.current,
						target: metrics.target,
						performance: metrics.current / metrics.target,
						trend: metrics.trend,
					}),
				),
				qualityMetrics: performance.productivity.codeQuality,
				teamSatisfaction: performance.collaboration.satisfaction,
				collaborationHealth: {
					crossTeamWork: performance.collaboration.crossTeamProjects,
					knowledgeSharing:
						performance.collaboration.knowledgeSharing.techTalks,
					mentoring: performance.collaboration.knowledgeSharing.mentoring,
				},
			},
		},
		analysisObjectives: [
			"Identify resource allocation inefficiencies",
			"Assess team composition balance (senior/mid/junior ratios)",
			"Evaluate project portfolio health and risk distribution",
			"Analyze performance trends and blockers",
			"Recommend organizational structure optimizations",
			"Suggest talent acquisition priorities",
		],
		contextualFactors: [
			"Department is in growth phase with 62% headcount increase planned",
			"Multiple high-priority projects running concurrently",
			"Budget utilization is high at 85%",
			"Average tenure is relatively low at 1.8 years",
			"One project is currently blocked, affecting timeline",
		],
	});

	// 3. Feed to AI - Mock AI call with realistic options
	const result = await callAI(analysisPrompt, {
		model: "claude-3-5-sonnet",
		temperature: 0.2,
		maxTokens: 2000,
		systemPrompt:
			"You are an expert organizational analyst specializing in engineering team optimization and resource allocation",
	});

	return {
		department: orgChart.name,
		analysis: result,
		timestamp: new Date().toISOString(),
		dataPoints: {
			teamsAnalyzed: orgChart.teams.length,
			projectsEvaluated: projects.length,
			totalMembers: orgChart.metrics.totalHeadcount,
			budgetScope: orgChart.metrics.totalBudget,
		},
	};
}

// Example usage:
// const orgAnalysis = await analyzeOrganizationalStructure("engineering")
// console.log(`Analysis for ${orgAnalysis.department} Department:`)
// console.log(orgAnalysis.analysis)
// console.log(`Data points: ${orgAnalysis.dataPoints.teamsAnalyzed} teams, ${orgAnalysis.dataPoints.projectsEvaluated} projects`)
// console.log(`Budget scope: $${orgAnalysis.dataPoints.budgetScope.toLocaleString()}`)

// Mock external services (replace with your actual services)
const orgService = {
	async getDepartment({ id }: { id: string }) {
		return {
			id,
			name: "Engineering",
			head: {
				name: "Sarah Chen",
				title: "VP of Engineering",
				experience: "8 years",
				previousRoles: ["Senior Director", "Engineering Manager"],
			},
			teams: [
				{
					name: "Backend Platform",
					lead: "Alex Rodriguez",
					members: [
						{
							name: "Maria Santos",
							role: "Senior Backend Engineer",
							seniority: "senior",
							joinDate: "2022-03-15",
						},
						{
							name: "James Kim",
							role: "Backend Engineer",
							seniority: "mid",
							joinDate: "2023-01-10",
						},
						{
							name: "Nina Patel",
							role: "Junior Backend Engineer",
							seniority: "junior",
							joinDate: "2023-09-01",
						},
					],
					technologies: ["Node.js", "PostgreSQL", "Redis", "Kubernetes"],
					budget: 450000,
					headcount: { current: 3, planned: 5 },
				},
				{
					name: "Frontend Experience",
					lead: "Elena Vasquez",
					members: [
						{
							name: "David Liu",
							role: "Senior Frontend Engineer",
							seniority: "senior",
							joinDate: "2021-11-20",
						},
						{
							name: "Anna Schmidt",
							role: "Frontend Engineer",
							seniority: "mid",
							joinDate: "2023-04-12",
						},
						{
							name: "Tom Wilson",
							role: "UI/UX Engineer",
							seniority: "mid",
							joinDate: "2022-08-30",
						},
					],
					technologies: ["React", "TypeScript", "Next.js", "Tailwind"],
					budget: 420000,
					headcount: { current: 3, planned: 4 },
				},
				{
					name: "DevOps & Infrastructure",
					lead: "Michael Johnson",
					members: [
						{
							name: "Lisa Zhang",
							role: "Senior DevOps Engineer",
							seniority: "senior",
							joinDate: "2020-05-15",
						},
						{
							name: "Carlos Ruiz",
							role: "Platform Engineer",
							seniority: "mid",
							joinDate: "2022-12-01",
						},
					],
					technologies: ["AWS", "Terraform", "Docker", "Prometheus"],
					budget: 380000,
					headcount: { current: 2, planned: 4 },
				},
			],
			metrics: {
				totalHeadcount: 8,
				plannedHeadcount: 13,
				averageTenure: "1.8 years",
				totalBudget: 1250000,
				utilizationRate: 0.85,
			},
		};
	},
};

const projectTracker = {
	async getActiveProjects({ _department }: { _department: string }) {
		return [
			{
				id: "proj-001",
				name: "API Gateway Migration",
				status: "in-progress",
				priority: "high",
				assignedTeam: "Backend Platform",
				assignedMembers: ["Maria Santos", "James Kim"],
				timeline: {
					start: "2024-01-15",
					deadline: "2024-03-30",
					progress: 0.65,
				},
				budget: { allocated: 120000, spent: 78000, remaining: 42000 },
				dependencies: ["Infrastructure upgrade", "Security review"],
				risks: [
					{
						description: "Third-party API deprecation",
						severity: "medium",
						mitigation: "Parallel development track",
					},
					{
						description: "Resource allocation conflict",
						severity: "low",
						mitigation: "Weekly sync meetings",
					},
				],
			},
			{
				id: "proj-002",
				name: "Mobile App Redesign",
				status: "planning",
				priority: "medium",
				assignedTeam: "Frontend Experience",
				assignedMembers: ["David Liu", "Tom Wilson"],
				timeline: {
					start: "2024-02-01",
					deadline: "2024-06-15",
					progress: 0.15,
				},
				budget: { allocated: 200000, spent: 25000, remaining: 175000 },
				dependencies: ["Design system completion", "User research"],
				risks: [
					{
						description: "Design approval delays",
						severity: "high",
						mitigation: "Stakeholder alignment sessions",
					},
				],
			},
			{
				id: "proj-003",
				name: "Kubernetes Migration",
				status: "blocked",
				priority: "high",
				assignedTeam: "DevOps & Infrastructure",
				assignedMembers: ["Lisa Zhang", "Carlos Ruiz"],
				timeline: {
					start: "2023-11-01",
					deadline: "2024-02-28",
					progress: 0.3,
				},
				budget: { allocated: 150000, spent: 95000, remaining: 55000 },
				dependencies: ["Security compliance review"],
				risks: [
					{
						description: "Legacy system compatibility",
						severity: "high",
						mitigation: "Gradual migration approach",
					},
					{
						description: "Budget overrun",
						severity: "medium",
						mitigation: "Resource reallocation from other projects",
					},
				],
			},
		];
	},
};

const performanceTracker = {
	async getTeamMetrics({ _department }: { _department: string }) {
		return {
			productivity: {
				velocity: {
					"Backend Platform": { current: 42, target: 45, trend: "improving" },
					"Frontend Experience": { current: 38, target: 40, trend: "stable" },
					"DevOps & Infrastructure": {
						current: 35,
						target: 38,
						trend: "declining",
					},
				},
				codeQuality: {
					testCoverage: 0.87,
					bugRate: 0.03,
					codeReviewTime: "2.3 hours",
					deploymentFrequency: "daily",
				},
			},
			collaboration: {
				crossTeamProjects: 3,
				knowledgeSharing: {
					techTalks: 8,
					documentation: "good",
					mentoring: "active",
				},
				satisfaction: {
					overallScore: 4.2,
					workLifeBalance: 4.0,
					careerGrowth: 3.8,
					teamDynamics: 4.5,
				},
			},
		};
	},
};

// Mock AI function (replace with your actual AI service)
async function callAI(
	_prompt: string,
	_options: {
		model?: string;
		temperature?: number;
		maxTokens?: number;
		systemPrompt?: string;
	},
) {
	// Simulate API delay
	await new Promise((resolve) => setTimeout(resolve, 900));

	return `## Engineering Department Organizational Analysis

### Executive Summary
The Engineering department shows strong fundamentals but faces scaling challenges typical of rapid growth phases. With a 62% planned headcount increase, strategic adjustments are needed to maintain team effectiveness and project delivery.

### 🎯 Key Findings

**Team Composition Balance:**
- Backend Platform: Well-balanced with 1:1:1 senior/mid/junior ratio
- Frontend Experience: Heavy on mid-level (2 mid, 1 senior) - needs senior reinforcement
- DevOps & Infrastructure: Top-heavy with senior expertise but understaffed overall

**Project Portfolio Health:**
- 2/3 high-priority projects (67%) indicates good strategic focus
- 1 blocked project (Kubernetes Migration) creates delivery risk
- Budget tracking is healthy with 65% average completion vs spend

**Performance Insights:**
- DevOps team velocity declining (35/38, 92% of target) - correlation with infrastructure blockers
- Backend team improving velocity (42/45, 93% of target) despite resource constraints
- Overall team satisfaction strong at 4.2/5, but career growth concerns at 3.8/5

### ⚠️ Critical Issues

1. **DevOps Bottleneck:** Only 2 engineers supporting entire infrastructure for 8+ person team
2. **Kubernetes Migration Blocking:** High-priority project stalled due to compliance review
3. **Frontend Leadership Gap:** Team lacks sufficient senior guidance for complex projects
4. **Budget Pressure:** 85% utilization leaves little room for unexpected needs

### 📋 Strategic Recommendations

**Immediate Actions (0-30 days):**
1. Unblock Kubernetes migration through dedicated compliance resource
2. Hire Senior Frontend Engineer to strengthen leadership
3. Add 1 additional DevOps engineer to alleviate bottleneck

**Short-term Optimization (30-90 days):**
4. Cross-train backend engineers in DevOps practices
5. Implement formal mentoring program to address career growth concerns
6. Establish project prioritization framework to manage concurrent high-priority work

**Long-term Structure (90+ days):**
7. Consider splitting Backend Platform team as it grows beyond 5 members
8. Create Architecture Review Board for better cross-team coordination
9. Implement rotation programs for knowledge sharing between teams

### 💰 Budget Allocation Recommendations
- Prioritize DevOps hiring (critical bottleneck): $180K
- Senior Frontend Engineer: $160K
- Defer some Backend growth until Q2 to manage burn rate
- Reserve 10% of remaining budget for compliance consultancy

### 📊 Success Metrics
- DevOps team velocity back to target (38) within 60 days
- Kubernetes migration completion by March 31st
- Frontend team velocity improvement to 42+ by Q2
- Overall team satisfaction maintained above 4.0

This analysis reveals a department with strong technical capabilities but needing structural adjustments to handle planned growth effectively. The focus should be on resolving bottlenecks before scaling further.`;
}

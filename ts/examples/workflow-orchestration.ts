import { llml } from "../src/index";

/**
 * Workflow Orchestration Example: CI/CD Pipeline Management
 *
 * Demonstrates how LLML transforms complex workflow state data into
 * structured prompts for AI-driven pipeline decisions. Shows the pattern:
 * Workflow State → LLML Transformation → Structured Decision Making
 */

export async function orchestrateDeploymentWorkflow(
	pipelineId: string,
	runId: string,
	environment: string,
): Promise<{
	pipelineId: string;
	runId: string;
	decision: {
		nextAction: string;
		reasoning: string;
		steps: {
			id: string;
			action: string;
			priority: string;
			estimatedDuration: string;
			dependencies: string[];
		}[];
		riskAssessment: {
			level: string;
			factors: string[];
			mitigations: string[];
		};
		monitoring: {
			metrics: string[];
			alertThresholds: {
				error_rate: number;
				latency_p95: number;
				throughput: number;
			};
			checkpointInterval: string;
		};
	};
	timestamp: string;
	executionContext: {
		stage: string;
		progress: number;
		metrics: {
			errorRate: number;
			latency: number;
			throughput: number;
			cpuUsage: number;
			memoryUsage: number;
		};
	};
}> {
	// 1. Fetch data from external services
	const pipeline = await workflowEngine.getDefinition({ id: pipelineId });
	const currentState = await executionTracker.getState({ runId });

	// 2. Transform with LLML - structure complex workflow context
	const workflowPrompt = llml({
		task: "analyze current workflow state and determine next actions",
		pipeline: {
			definition: pipeline,
			currentExecution: {
				runId: currentState.runId,
				stage: currentState.currentStage,
				progress: `${Math.round(currentState.progress * 100)}%`,
				duration: `${Math.round((Date.now() - new Date(currentState.startTime).getTime()) / 60000)}m elapsed`,
			},
		},
		environment: {
			target: environment,
			constraints: [
				"Zero downtime requirement",
				"Must maintain 99.9% availability",
				"Rollback within 3 minutes if issues detected",
			],
		},
		currentMetrics: {
			performance: {
				errorRate: currentState.metrics.errorRate,
				avgLatency: `${currentState.metrics.latency}ms`,
				throughput: `${currentState.metrics.throughput} req/min`,
			},
			infrastructure: {
				cpuUsage: `${Math.round(currentState.metrics.cpuUsage * 100)}%`,
				memoryUsage: `${Math.round(currentState.metrics.memoryUsage * 100)}%`,
			},
			healthStatus: currentState.healthChecks,
		},
		context: {
			recentHistory: currentState.previousRuns.map((run) => ({
				outcome: run.status,
				duration: run.duration,
				issues: run.reason || "none",
			})),
			businessImpact: environment === "production" ? "high" : "medium",
			timeOfDay: new Date().getHours() < 17 ? "business hours" : "after hours",
		},
		decisionCriteria: [
			"Prioritize system stability over speed",
			"Consider business impact and timing",
			"Ensure proper monitoring is in place",
			"Have clear rollback procedures ready",
		],
	});

	// 3. Feed to AI - Mock AI call with realistic options
	const decision = await callAI(workflowPrompt, {
		model: "gpt-4",
		temperature: 0.2,
		responseFormat: "json",
		systemPrompt:
			"You are an expert DevOps engineer specializing in CI/CD pipeline management and risk assessment",
	});

	return {
		pipelineId,
		runId,
		decision,
		timestamp: new Date().toISOString(),
		executionContext: {
			stage: currentState.currentStage,
			progress: currentState.progress,
			metrics: currentState.metrics,
		},
	};
}

// Example usage:
// const orchestration = await orchestrateDeploymentWorkflow("deploy-prod-pipeline", "run-456", "production")
// console.log(`Next action: ${orchestration.decision.nextAction}`)
// console.log(`Risk level: ${orchestration.decision.riskAssessment.level}`)
// console.log(`Steps: ${orchestration.decision.steps.length}`)

// Mock external services (replace with your actual services)
const workflowEngine = {
	async getDefinition({ id }: { id: string }) {
		return {
			id,
			name: "Production Deployment Pipeline",
			stages: [
				{ name: "pre-checks", status: "completed", duration: "2m" },
				{ name: "build", status: "completed", duration: "8m" },
				{ name: "test", status: "completed", duration: "12m" },
				{ name: "staging-deploy", status: "completed", duration: "5m" },
				{ name: "staging-validation", status: "in-progress", duration: "15m" },
				{ name: "prod-deploy", status: "pending", duration: "10m" },
				{ name: "health-check", status: "pending", duration: "5m" },
			],
			rollbackStrategy: "automated",
			maxRollbackTime: "3 minutes",
		};
	},
};

const executionTracker = {
	async getState({ runId }: { runId: string }) {
		return {
			runId,
			startTime: "2024-01-15T14:30:00Z",
			currentStage: "staging-validation",
			progress: 0.71,
			metrics: {
				errorRate: 0.002,
				latency: 245,
				throughput: 1200,
				cpuUsage: 0.65,
				memoryUsage: 0.78,
			},
			healthChecks: {
				database: "healthy",
				cache: "healthy",
				externalServices: "degraded",
				loadBalancer: "healthy",
			},
			previousRuns: [
				{ id: "run-123", status: "success", duration: "24m" },
				{
					id: "run-122",
					status: "rollback",
					duration: "18m",
					reason: "high error rate",
				},
			],
		};
	},
};

// Mock AI function (replace with your actual AI service)
async function callAI(
	_prompt: string,
	_options: {
		model?: string;
		temperature?: number;
		responseFormat?: string;
		systemPrompt?: string;
	},
) {
	// Simulate API delay
	await new Promise((resolve) => setTimeout(resolve, 700));

	return {
		nextAction: "continue",
		reasoning:
			"Current metrics are within acceptable thresholds. Error rate at 0.2% is well below 1% limit, latency at 245ms is acceptable, and most health checks are green. The degraded external services status is concerning but not blocking since it's not critical path.",
		steps: [
			{
				id: "validate-external-deps",
				action: "Investigate degraded external services status",
				priority: "high",
				estimatedDuration: "5 minutes",
				dependencies: [],
			},
			{
				id: "complete-staging-validation",
				action: "Finish staging validation with enhanced monitoring",
				priority: "high",
				estimatedDuration: "10 minutes",
				dependencies: ["validate-external-deps"],
			},
			{
				id: "proceed-to-production",
				action: "Begin production deployment using blue-green strategy",
				priority: "medium",
				estimatedDuration: "15 minutes",
				dependencies: ["complete-staging-validation"],
			},
		],
		riskAssessment: {
			level: "medium",
			factors: [
				"External services degradation could affect production",
				"Recent rollback in run-122 indicates system sensitivity",
				"High business impact due to production environment",
			],
			mitigations: [
				"Enhanced monitoring during production deployment",
				"Automated rollback triggers set at stricter thresholds",
				"On-call team standing by for immediate response",
			],
		},
		monitoring: {
			metrics: [
				"error_rate",
				"latency_p95",
				"throughput",
				"external_service_health",
			],
			alertThresholds: {
				error_rate: 0.005,
				latency_p95: 400,
				throughput: 1000,
			},
			checkpointInterval: "2 minutes",
		},
	};
}

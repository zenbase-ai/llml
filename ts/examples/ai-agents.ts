import { llml } from "../src/index"

/**
 * AI Agents Example: Multi-Agent System Coordination
 *
 * Demonstrates how LLML transforms complex agent coordination data into
 * structured prompts for AI systems. Shows the pattern:
 * Data Fetching → LLML Transformation → AI Generation
 */

export async function coordinateMultiAgentDeployment(applicationId: string, targetEnvironment: string) {
  // 1. Fetch data from external services
  const agents = await agentRegistry.getAgents({ _team: "devops" })
  const capabilities = await capabilityService.getCapabilities({ _domain: "deployment" })

  // 2. Transform with LLML - this is where the magic happens
  const agentPrompt = llml({
    task: "coordinate multi-agent deployment workflow",
    application: {
      id: applicationId,
      targetEnvironment,
      requiresDowntime: false,
      criticalityLevel: "high",
    },
    availableAgents: agents.map(agent => ({
      id: agent.id,
      role: agent.role,
      expertise: agent.skills,
      availability: agent.status === "active" ? 1 - agent.currentLoad : 0,
    })),
    systemCapabilities: capabilities,
    constraints: [
      "Minimize deployment risk",
      "Ensure zero-downtime deployment",
      "Validate security compliance",
      "Monitor performance metrics",
    ],
    coordinationRules: [
      "Deployment coordinator leads the workflow",
      "Security validator must approve before deployment",
      "Monitoring specialist tracks all metrics during deployment",
      "Any agent can trigger rollback if issues detected",
    ],
  })

  // 3. Feed to AI - Mock AI call with realistic options
  const result = await callAI(agentPrompt, {
    model: "gpt-4",
    temperature: 0.3,
    maxTokens: 1000,
    systemPrompt: "You are an expert DevOps coordinator specializing in multi-agent deployment workflows",
  })

  return {
    coordinationPlan: result,
    agentsInvolved: agents.length,
    estimatedDuration: "15-30 minutes",
  }
}

// Example usage:
// const deployment = await coordinateMultiAgentDeployment("web-app-v2", "production")
// console.log(deployment.coordinationPlan)

// Mock external services (replace with your actual services)
const agentRegistry = {
  async getAgents({ _team }: { _team: string }) {
    return [
      {
        id: "deploy-agent-1",
        role: "deployment-coordinator",
        skills: ["kubernetes", "docker", "helm"],
        status: "active",
        currentLoad: 0.3,
      },
      {
        id: "monitor-agent-2",
        role: "monitoring-specialist",
        skills: ["prometheus", "grafana", "alerting"],
        status: "active",
        currentLoad: 0.7,
      },
      {
        id: "security-agent-3",
        role: "security-validator",
        skills: ["vulnerability-scanning", "compliance", "access-control"],
        status: "active",
        currentLoad: 0.2,
      },
    ]
  },
}

const capabilityService = {
  async getCapabilities({ _domain }: { _domain: string }) {
    return {
      deployment: {
        strategies: ["blue-green", "canary", "rolling"],
        rollbackTime: "< 5 minutes",
        maxConcurrentDeployments: 3,
      },
      monitoring: {
        metrics: ["cpu", "memory", "network", "errors"],
        alertThresholds: { errorRate: 0.01, latency: "500ms" },
        retentionPeriod: "30 days",
      },
      security: {
        scanTypes: ["sast", "dast", "dependency"],
        complianceStandards: ["soc2", "pci-dss"],
        accessPolicies: ["rbac", "zero-trust"],
      },
    }
  },
}

// Mock AI function (replace with your actual AI service)
async function callAI(
  _prompt: string,
  _options: {
    model?: string
    temperature?: number
    maxTokens?: number
    systemPrompt?: string
  },
) {
  // Simulate API delay
  await new Promise(resolve => setTimeout(resolve, 500))

  return `## Multi-Agent Deployment Coordination Plan

### Agent Assignment
- **deploy-agent-1** (deployment-coordinator): Lead the deployment workflow
- **security-agent-3** (security-validator): Pre-deployment security validation  
- **monitor-agent-2** (monitoring-specialist): Real-time monitoring during deployment

### Execution Strategy
1. **Pre-deployment Phase** (5 minutes)
   - Security agent performs vulnerability scan and compliance check
   - Deployment agent validates infrastructure readiness
   - Monitor agent establishes baseline metrics

2. **Deployment Phase** (15 minutes)
   - Deploy using blue-green strategy to minimize risk
   - Monitor agent tracks error rates, latency, and throughput
   - Security agent monitors for any suspicious activity

3. **Validation Phase** (10 minutes)
   - Run automated health checks across all services
   - Validate business-critical functionality
   - Confirm performance metrics within acceptable thresholds

### Risk Mitigation
- Automated rollback triggers if error rate > 1% or latency > 500ms
- Security agent has veto power over deployment if vulnerabilities detected
- Monitor agent maintains 2-minute alert response time

### Success Criteria
- Zero downtime achieved
- All security validations passed
- Performance metrics stable within 5% of baseline
- All agents report "green" status

This coordinated approach leverages each agent's specialized capabilities while maintaining clear accountability and rapid response protocols.`
}

/**
 * Integration tests for LLML TypeScript examples
 */

import { describe, expect, it } from "vitest"

import { coordinateMultiAgentDeployment } from "../examples/ai-agents"
import { validateEnvironmentConfig } from "../examples/config-management"
import { analyzeOrganizationalStructure } from "../examples/data-structures"
import { queryDocumentation } from "../examples/rag-context"
import { orchestrateDeploymentWorkflow } from "../examples/workflow-orchestration"

// Comprehensive sanity check for all TypeScript examples
// Ensures the mocked services + LLML transformation logic stay in sync

describe("LLML TypeScript Examples", () => {
  it("AI Agents Example", async () => {
    const result = await coordinateMultiAgentDeployment("web-app-v2", "production")
    expect(result.agentsInvolved).toBe(3)
    expect(result.estimatedDuration).toBe("15-30 minutes")
    expect(result.coordinationPlan).toContain("Multi-Agent Deployment Coordination Plan")
  })

  it("RAG Context Example", async () => {
    const result = await queryDocumentation("How do I authenticate with your API?")
    expect(result.documentsSearched).toBe(3)
    expect(result.analysis.confidence).toBeCloseTo(0.95, 2)
    expect(result.analysis.citations.length).toBe(1)
  })

  it("Workflow Orchestration Example", async () => {
    const result = await orchestrateDeploymentWorkflow("deploy-prod-pipeline", "run-456", "production")
    expect(result.decision.nextAction).toBe("continue")
    expect(result.decision.riskAssessment.level).toBe("medium")
    expect(result.decision.steps.length).toBe(3)
  })

  it("Config Management Example", async () => {
    const result = await validateEnvironmentConfig("production", "web-api")
    expect(result.summary.secretsAnalyzed).toBe(4)
    expect(result.summary.complianceStandards).toBe(3)
    expect(result.configurationAnalysis).toContain("Configuration Health Assessment")
  })

  it("Data Structures Example", async () => {
    const result = await analyzeOrganizationalStructure("engineering")
    expect(result.dataPoints.teamsAnalyzed).toBe(3)
    expect(result.dataPoints.totalMembers).toBe(8)
    expect(result.dataPoints.budgetScope).toBe(1250000)
    expect(result.analysis).toContain("Organizational Analysis")
  })
})

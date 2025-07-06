import { llml } from "../src/index";

/**
 * Configuration Management Example: Environment-Aware Config Validation
 *
 * Demonstrates how LLML transforms complex configuration data into
 * structured prompts for AI-driven configuration validation and optimization.
 * Shows the pattern: Config Sources → LLML Transformation → AI Analysis
 */

export async function validateEnvironmentConfig(
	environment: string,
	applicationName: string,
): Promise<{
	application: string;
	environment: string;
	configurationAnalysis: string;
	timestamp: string;
	summary: {
		secretsAnalyzed: number;
		complianceStandards: number;
		configurationAreas: number;
	};
}> {
	// 1. Fetch data from external services
	const envConfig = await configService.getEnvConfig({ env: environment });
	const secrets = await secretsManager.getSecrets({
		_namespace: applicationName,
	});
	const compliance = await complianceService.getRequirements({
		env: environment,
	});

	// 2. Transform with LLML - structure configuration analysis context
	const configPrompt = llml({
		task: "analyze and validate environment configuration for security, performance, and compliance",
		application: {
			name: applicationName,
			targetEnvironment: environment,
			deploymentTimestamp: new Date().toISOString(),
		},
		configuration: {
			infrastructure: {
				scaling: envConfig.scaling,
				database: {
					...envConfig.database,
					// Never include actual connection strings in prompts
					connectionSecure: envConfig.database.ssl
						? "SSL enabled"
						: "SSL disabled",
				},
				networking: envConfig.networking,
			},
			features: envConfig.features,
			security: {
				secretsCount: secrets.length,
				secretsHealth: secrets.map((s) => ({
					type: s.type,
					rotationPolicy: s.rotation,
					daysSinceRotation: Math.floor(
						(Date.now() - new Date(s.lastRotated).getTime()) /
							(1000 * 60 * 60 * 24),
					),
					status: s.rotation === "manual" ? "needs-review" : "automated",
				})),
			},
		},
		compliance: {
			requiredStandards: compliance.standards,
			encryptionRequirements: compliance.requirements.encryption,
			loggingRequirements: compliance.requirements.logging,
			accessControlRequirements: compliance.requirements.access,
		},
		validationCriteria: [
			"Ensure security best practices are followed",
			"Verify compliance with required standards",
			"Check for performance optimization opportunities",
			"Identify potential configuration drift or issues",
			"Validate secret rotation policies are appropriate",
		],
		outputFormat: [
			"Provide an overall configuration health score (1-10)",
			"List any security vulnerabilities or concerns",
			"Identify compliance gaps if any exist",
			"Suggest performance optimizations",
			"Recommend action items with priority levels",
		],
	});

	// 3. Feed to AI - Mock AI call with realistic options
	const analysis = await callAI(configPrompt, {
		model: "gpt-4",
		temperature: 0.1,
		maxTokens: 2000,
		systemPrompt:
			"You are an expert DevOps engineer specializing in configuration management, security, and compliance. Provide detailed, actionable analysis.",
	});

	return {
		application: applicationName,
		environment,
		configurationAnalysis: analysis,
		timestamp: new Date().toISOString(),
		summary: {
			secretsAnalyzed: secrets.length,
			complianceStandards: compliance.standards.length,
			configurationAreas: Object.keys(envConfig).length,
		},
	};
}

// Example usage:
// const validation = await validateEnvironmentConfig("production", "web-api")
// console.log(`Configuration Analysis for ${validation.application}:`)
// console.log(validation.configurationAnalysis)
// console.log(`Secrets analyzed: ${validation.summary.secretsAnalyzed}`)
// console.log(`Compliance standards: ${validation.summary.complianceStandards}`)

// Mock external services (replace with your actual services)
const configService = {
	async getEnvConfig({ env }: { env: string }) {
		return {
			environment: env,
			region: "us-east-1",
			scaling: {
				minInstances: env === "production" ? 3 : 1,
				maxInstances: env === "production" ? 20 : 5,
				targetCPU: 70,
				targetMemory: 80,
			},
			database: {
				host:
					env === "production" ? "prod-db-cluster.internal" : "dev-db.internal",
				port: 5432,
				poolSize: env === "production" ? 50 : 10,
				timeout: env === "production" ? 30000 : 10000,
				ssl: env === "production",
				backupEnabled: env === "production",
			},
			features: {
				rateLimiting: env === "production",
				debugging: env !== "production",
				telemetry: true,
				caching: env === "production" ? "redis" : "memory",
			},
			networking: {
				allowedOrigins:
					env === "production"
						? ["https://app.example.com", "https://api.example.com"]
						: ["http://localhost:3000", "http://localhost:8080"],
				loadBalancer: env === "production" ? "application" : "none",
				cdn: env === "production",
			},
		};
	},
};

const secretsManager = {
	async getSecrets({ _namespace }: { _namespace: string }) {
		return [
			{
				key: "DATABASE_PASSWORD",
				masked: "prod_****_2024",
				type: "password",
				rotation: "monthly",
				lastRotated: "2024-01-01T00:00:00Z",
			},
			{
				key: "JWT_SECRET",
				masked: "jwt_****_key",
				type: "signing-key",
				rotation: "quarterly",
				lastRotated: "2023-10-01T00:00:00Z",
			},
			{
				key: "STRIPE_API_KEY",
				masked: "sk_live_****",
				type: "api-key",
				rotation: "yearly",
				lastRotated: "2023-06-01T00:00:00Z",
			},
			{
				key: "SENDGRID_API_KEY",
				masked: "SG.****.*****",
				type: "api-key",
				rotation: "manual",
				lastRotated: "2023-12-15T00:00:00Z",
			},
		];
	},
};

const complianceService = {
	async getRequirements({ env }: { env: string }) {
		return {
			standards: env === "production" ? ["SOC2", "PCI-DSS", "GDPR"] : ["GDPR"],
			requirements: {
				encryption: {
					atRest: env === "production",
					inTransit: true,
					keyRotation: env === "production" ? "quarterly" : "yearly",
				},
				logging: {
					auditTrail: env === "production",
					retention: env === "production" ? "7 years" : "1 year",
					level: env === "production" ? "INFO" : "DEBUG",
				},
				access: {
					mfa: env === "production",
					rbac: true,
					sessionTimeout: env === "production" ? "4 hours" : "8 hours",
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
	await new Promise((resolve) => setTimeout(resolve, 800));

	return `## Configuration Health Assessment

### Overall Score: 8.5/10
The production configuration for web-api shows strong security posture with room for optimization.

### ✅ Security Strengths
- SSL encryption properly enabled for database connections
- Production-appropriate scaling configuration (3-20 instances)
- Rate limiting enabled for production environment
- RBAC access controls in place
- All critical secrets properly masked and managed

### ⚠️ Areas of Concern

**Secret Rotation Issues:**
- JWT_SECRET last rotated 3 months ago (quarterly policy)
- SENDGRID_API_KEY on manual rotation (outdated by 15 days)

**Performance Optimizations:**
- Database pool size (50) may be excessive for current load
- Memory target (80%) leaves little headroom for spikes

### 🔴 Compliance Gaps
- SOC2 audit trail logging not explicitly configured
- PCI-DSS encryption at rest missing for database
- GDPR data retention policies not defined in configuration

### 📋 Recommended Actions

**High Priority:**
1. Rotate SENDGRID_API_KEY immediately (overdue)
2. Enable database encryption at rest for PCI compliance
3. Configure audit trail logging for SOC2 requirements

**Medium Priority:**
4. Review database pool size optimization (reduce to 30-35)
5. Implement automated JWT secret rotation
6. Add GDPR data retention configuration

**Low Priority:**
7. Adjust memory target to 75% for better headroom
8. Review CDN configuration for performance gains

### Implementation Timeline
- **Immediate (1-2 days):** Secret rotation and encryption
- **Short-term (1 week):** Audit logging and compliance configs
- **Medium-term (2-4 weeks):** Performance optimizations

This configuration follows most best practices but requires immediate attention to secret rotation and compliance gaps before production deployment.`;
}

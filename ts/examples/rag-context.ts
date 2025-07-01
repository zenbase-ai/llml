import { llml } from "../src/index"

/**
 * RAG Context Example: Document Retrieval + Structured Analysis
 *
 * Demonstrates how LLML transforms RAG document retrieval data into
 * structured prompts for AI analysis. Shows the pattern:
 * Vector Search → LLML Transformation → Structured AI Response
 */

export async function queryDocumentation(userQuery: string) {
  // 1. Fetch data from external services
  const docs = await vectorDB.similaritySearch(userQuery, { limit: 5 })
  const metadata = await documentStore.getMetadata(docs.map(d => d.id))

  // 2. Transform with LLML - create structured context for AI
  const ragPrompt = llml({
    system: "You are a helpful documentation assistant that provides accurate, well-cited answers",
    instructions: [
      "Answer the user's question based only on the provided documentation",
      "Always cite specific documents when making claims",
      "If information is missing or unclear, explicitly state this",
      "Provide confidence scores based on document relevance and completeness",
    ],
    context: {
      query: userQuery,
      retrievedDocuments: docs.map(doc => ({
        title: doc.metadata.source,
        content: doc.content,
        relevanceScore: doc.score,
        lastUpdated: doc.metadata.lastUpdated,
      })),
      corpus: {
        totalDocuments: metadata.totalDocuments,
        lastIndexed: metadata.lastIndexed,
        coverage: metadata.documentTypes,
      },
    },
    outputRequirements: [
      "Provide a direct answer with confidence score",
      "Include relevant citations with quotes",
      "Suggest follow-up questions if appropriate",
      "Identify any knowledge gaps in the available documents",
    ],
  })

  // 3. Feed to AI - Mock AI call with realistic options
  const analysis = await callAI(ragPrompt, {
    model: "gpt-4o",
    temperature: 0.1,
    responseFormat: "json",
    systemPrompt: "You are an expert documentation analyst. Provide structured, well-cited responses.",
  })

  return {
    query: userQuery,
    analysis,
    documentsSearched: docs.length,
    processingTime: Date.now(),
  }
}

// Example usage:
// const result = await queryDocumentation("How do I authenticate with your API?")
// console.log(`Answer: ${result.analysis.answer}`)
// console.log(`Confidence: ${result.analysis.confidence}`)
// console.log(`Citations: ${result.analysis.citations.length}`)

// Mock external services (replace with your actual services)
const vectorDB = {
  async similaritySearch(_query: string, _options: { limit: number }) {
    return [
      {
        id: "doc-1",
        content:
          "OAuth 2.0 is the industry standard for authorization. Our API uses OAuth 2.0 with PKCE for secure authentication. Clients must register and obtain client credentials.",
        score: 0.95,
        metadata: { source: "api-auth-guide.md", lastUpdated: "2024-01-15" },
      },
      {
        id: "doc-2",
        content:
          "Rate limiting is enforced at 1000 requests per hour per API key. Premium accounts get 10,000 requests per hour. Use the X-RateLimit headers to track usage.",
        score: 0.82,
        metadata: { source: "rate-limits.md", lastUpdated: "2024-01-10" },
      },
      {
        id: "doc-3",
        content:
          "API errors follow REST conventions. 401 means unauthorized, 403 means forbidden, 429 means rate limited. All errors include a detailed message and error code.",
        score: 0.78,
        metadata: { source: "error-handling.md", lastUpdated: "2024-01-08" },
      },
    ]
  },
}

const documentStore = {
  async getMetadata(_docIds: string[]) {
    return {
      totalDocuments: 156,
      lastIndexed: "2024-01-15T10:30:00Z",
      documentTypes: ["guides", "references", "tutorials"],
      averageAge: "12 days",
    }
  },
}

// Mock AI function (replace with your actual AI service)
async function callAI(
  _prompt: string,
  _options: {
    model?: string
    temperature?: number
    responseFormat?: string
    systemPrompt?: string
  },
) {
  // Simulate API delay
  await new Promise(resolve => setTimeout(resolve, 600))

  return {
    answer:
      "To authenticate with our API, you need to use OAuth 2.0 with PKCE (Proof Key for Code Exchange). First, register your application to obtain client credentials, then follow the OAuth 2.0 authorization flow to get access tokens.",
    confidence: 0.95,
    citations: [
      {
        documentTitle: "api-auth-guide.md",
        relevantQuote:
          "Our API uses OAuth 2.0 with PKCE for secure authentication. Clients must register and obtain client credentials.",
        relevanceScore: 0.95,
      },
    ],
    followupQuestions: [
      "How do I register my application for OAuth credentials?",
      "What scopes are available for API access?",
      "How do I refresh expired access tokens?",
    ],
    knowledgeGaps: [
      "Specific OAuth endpoint URLs not found in retrieved documents",
      "Token expiration times not clearly specified",
    ],
  }
}

# Manga Wiki

## Product Requirements Document — v0.1

**Status:** Draft / Product Foundation
**Document Type:** Product Requirements Document (PRD)
**Product Codename:** Manga Wiki
**Version:** 0.1
**Date:** 2026-09-07
**Role of this document:** Product Constitution / North Star

---

# 1. Executive Summary

## 1.1 Product

**Manga Wiki** is a domain-specific, evidence-grounded knowledge engine designed to answer deep, complex, multi-hop questions about manga.

The product combines:

* Wiki-style structured knowledge
* Hybrid RAG
* Knowledge Graph
* Entity-aware retrieval
* Event-aware retrieval
* Temporal reasoning
* Agentic retrieval
* Knowledge compilation
* Evidence provenance
* Source authority
* Canon awareness
* Contradiction detection
* Entity resolution
* Spoiler-aware retrieval
* Persistent knowledge
* Human-curated Wiki pages

The primary user interface may include:

* conversational AI
* Wiki pages
* search
* character pages
* chapter pages
* timelines
* relationship graphs
* source/evidence views

However, the product is **not fundamentally a chatbot**.

The core product is the underlying **Manga Knowledge Engine**.

---

# 2. Product Vision

## 2.1 Vision Statement

> Build the most reliable and deeply reasoned open knowledge system for manga, capable of answering questions that ordinary keyword search, vector RAG, or generic AI assistants cannot answer reliably.

Manga Wiki should allow users to explore manga knowledge as if they had:

* a comprehensive encyclopedia,
* a research assistant,
* a timeline engine,
* a relationship graph,
* and a domain-specialized reasoning agent

combined into one system.

---

# 3. Product Thesis

Traditional RAG generally follows:

```
User Question
    ↓
Retrieve Chunks
    ↓
LLM
    ↓
Answer
```

This architecture is insufficient for deep manga questions because manga knowledge is highly relational and temporal.

A manga question may depend on:

* multiple characters,
* multiple chapters,
* multiple events,
* chronology,
* aliases,
* organizations,
* locations,
* objects,
* abilities,
* previous events,
* later events,
* source authority,
* canon status,
* and contradictions between sources.

Manga Wiki therefore follows a different thesis:

```
Raw Sources
    ↓
Knowledge Extraction
    ↓
Entity Resolution
    ↓
Event Extraction
    ↓
Relationship Extraction
    ↓
Knowledge Compilation
    ↓
Wiki + Knowledge Graph + Evidence Store
    ↓
Multi-Strategy Retrieval
    ↓
Agentic Reasoning
    ↓
Evidence Verification
    ↓
Grounded Answer
```

The system should retrieve and reason over **knowledge**, not merely retrieve similar text chunks.

---

# 4. Core Product Principle

## 4.1 Wiki is not merely the UI

The Wiki should not be treated as a frontend layer placed on top of a vector database.

The Wiki represents a persistent, navigable, human-readable representation of the knowledge accumulated by the system.

Therefore:

```
Wiki
  =
Knowledge Representation
```

while:

```
RAG
  =
Retrieval Mechanism
```

and:

```
Agent
  =
Reasoning / Orchestration Mechanism
```

and:

```
Evidence
  =
Grounding / Verification Mechanism
```

---

# 5. Product Goals

## 5.1 Primary Goals

Manga Wiki must:

1. Answer manga questions with high factual accuracy.
2. Support questions requiring multiple retrieval steps.
3. Understand relationships between entities.
4. Understand events and chronology.
5. Distinguish canon from speculation.
6. Provide evidence for important factual claims.
7. Respect user spoiler boundaries.
8. Resolve aliases and entity ambiguity.
9. Combine structured and unstructured knowledge.
10. Maintain persistent knowledge rather than reconstructing everything for every query.
11. Support human inspection and correction of knowledge.
12. Make reasoning traceable without exposing private chain-of-thought.
13. Improve retrieval quality through domain-specific techniques.
14. Support multiple manga independently while allowing future cross-manga reasoning.
15. Remain modular enough that individual retrieval techniques can be evaluated and replaced.

---

# 6. Non-Goals

The following are explicitly not core goals of v1:

## 6.1 Manga reader

Manga Wiki is not intended to become a manga reading platform.

It should not redistribute unauthorized manga scans.

## 6.2 Generic AI assistant

The product should not attempt to become a general-purpose ChatGPT competitor.

Its intelligence should be optimized for manga knowledge.

## 6.3 Social network

Community features may exist later, but social interaction is not a primary product goal.

## 6.4 Image-generation platform

Image generation is not a core capability.

## 6.5 Blind scraping system

The product should not assume that "more scraped data" automatically means better knowledge.

Source quality and provenance are more important than raw volume.

---

# 7. Target Users

## 7.1 Manga Reader

Wants fast and accurate answers.

Examples:

* "Who is this character?"
* "When did X happen?"
* "What chapter introduced Y?"

## 7.2 Deep Fan

Wants detailed explanations.

Examples:

* "Explain the relationship between X and Y."
* "List every encounter between X and Y."
* "How did X's ideology change?"

## 7.3 Lore Researcher

Wants evidence-backed synthesis.

Examples:

* "What evidence supports this theory?"
* "What are all known connections between these characters?"
* "What events led to the current political structure?"

## 7.4 New Reader

Needs spoiler protection.

Example:

> "I have read up to chapter 400. Explain this character without spoilers."

## 7.5 Power / Battle Analyst

Wants structured comparison.

Examples:

* "Who defeated X?"
* "What abilities counter Y?"
* "How did X's combat ability evolve?"

## 7.6 Content Creator / Researcher

Needs citations and source traceability.

Examples:

* "Give me every chapter supporting this claim."
* "Build a timeline of these events."
* "Compare the manga with official supplementary material."

---

# 8. Core User Experience

The product should expose several complementary interfaces.

## 8.1 AI Chat

Primary natural-language interface.

Example:

> "Why did Shanks give Luffy the straw hat?"

The answer should contain:

* concise answer
* relevant context
* evidence
* chapter references
* confidence/qualification where appropriate
* spoiler boundary

---

## 8.2 Wiki

Every important entity should have a structured page.

Examples:

```
/manga/one-piece

/character/monkey-d-luffy

/character/shanks

/chapter/1

/arc/east-blue

/event/luffy-meets-shanks
```

---

## 8.3 Search

Search should support:

* entity search
* semantic search
* exact search
* chapter search
* event search
* relationship search
* source search

---

## 8.4 Knowledge Graph

Users should eventually be able to inspect relationships.

Example:

```
Shanks
   │
   ├── mentors → Luffy
   ├── knew → Roger
   ├── belongs_to → Red Hair Pirates
   └── possesses_history_with → Straw Hat
```

---

## 8.5 Timeline

Manga-specific chronological navigation.

The system should distinguish:

* publication chronology
* story chronology
* flashback chronology
* inferred chronology

---

# 9. Manga Domain Model

Manga Wiki should develop a domain ontology rather than treating all content as generic documents.

Core entities should include:

```
Manga
Series
Author
Artist
Publisher
Magazine
Volume
Chapter
Arc
Character
Organization
Faction
Location
Item
Ability
Event
Relationship
Quote
Theme
Concept
Species
Race
Family
Timeline
Source
Evidence
Theory
```

The ontology should remain extensible.

---

# 10. Character Model

A character should support:

* canonical name
* aliases
* translations
* titles
* affiliations
* relationships
* appearances
* abilities
* locations
* events
* first appearance
* latest known appearance
* status
* timeline
* evidence
* source authority
* canon status

Example conceptual structure:

```
Character
  ├── identity
  ├── aliases
  ├── affiliations
  ├── relationships
  ├── appearances
  ├── abilities
  ├── events
  ├── locations
  ├── timeline
  └── evidence
```

---

# 11. Event Model

Events are first-class objects.

This is critical.

Instead of storing only:

> "X defeated Y."

the system should conceptually represent:

```
Event
  type: Combat
  participants:
      X
      Y
  location:
      Z
  chapter:
      N
  outcome:
      X defeated Y
  preceding_events:
      ...
  following_events:
      ...
  evidence:
      ...
```

This allows temporal and causal questions.

---

# 12. Relationship Model

Relationships should be typed.

Examples:

```
knows
meets
mentors
hates
loves
fights
defeats
kills
protects
betrays
belongs_to
leads
serves
inherits
possesses
visits
lives_in
originates_from
appears_in
participates_in
```

Relationships may also be:

* directional
* temporal
* conditional
* disputed
* inferred
* explicitly confirmed

---

# 13. Temporal Model

Temporal reasoning is a first-class requirement.

The system should distinguish:

## 13.1 Publication time

When a chapter was published.

## 13.2 Story time

When an event occurred in the manga's internal chronology.

## 13.3 Relative time

Relationships such as:

* before
* after
* during
* immediately before
* shortly after
* years before
* years later

## 13.4 Uncertain time

Some events cannot be precisely dated.

The system must be able to represent uncertainty rather than invent dates.

---

# 14. Canon Model

Manga Wiki must not treat all sources as equally authoritative.

Potential source classes:

```
Primary manga
Official publication
Author statement
Official interview
Official databook
Official guide
Publisher material
Anime adaptation
Movie
Spin-off
Fan wiki
Community discussion
Fan theory
```

Each fact should conceptually have:

```
source_type
authority
canon_status
confidence
evidence
```

The system should be able to answer:

> "Is this actually canon?"

rather than merely:

> "Did some website say this?"

---

# 15. Evidence Model

Important claims should be traceable.

Conceptual structure:

```
Claim
   ↓
Evidence
   ↓
Source
   ↓
Document
   ↓
Chapter / Section / Page
```

An answer should ideally allow the user to inspect:

```
Claim
  ↓
Why do we believe this?
  ↓
Evidence
  ↓
Source
```

---

# 16. Evidence Grounding Requirements

The system SHOULD NOT present uncertain information as established fact.

It should distinguish:

### Confirmed

Directly supported by authoritative evidence.

### Strongly implied

Not directly stated but supported by multiple pieces of evidence.

### Possible

Plausible but insufficiently established.

### Speculative

Primarily theory or interpretation.

### Contradicted

Conflicts with stronger evidence.

---

# 17. Entity Resolution

Entity resolution is critical.

The system must understand that:

```
Monkey D. Luffy
Luffy
Straw Hat Luffy
Straw Hat
```

may refer to the same entity depending on context.

It must also handle:

* translations
* romanization differences
* alternate spellings
* nicknames
* titles
* aliases
* localization differences

Entity resolution should occur before high-level reasoning whenever possible.

---

# 18. Entity Deduplication

Knowledge extraction may repeatedly produce:

```
Shanks
Red Hair Shanks
Akagami
Red-Haired Shanks
```

These should not automatically become separate entities.

The system should maintain canonical entities with aliases.

---

# 19. Knowledge Compilation

Manga Wiki should progressively transform raw sources into structured knowledge.

Conceptual pipeline:

```
Raw Source
   ↓
Parsing
   ↓
Chunking / Segmentation
   ↓
Entity Extraction
   ↓
Entity Resolution
   ↓
Event Extraction
   ↓
Relationship Extraction
   ↓
Temporal Extraction
   ↓
Canon Classification
   ↓
Evidence Linking
   ↓
Knowledge Graph
   ↓
Wiki Pages
   ↓
Retrieval Indexes
```

The Wiki is therefore a **compiled knowledge layer**.

---

# 20. RAG Strategy

Manga Wiki should not depend on one retrieval strategy.

The system should support multiple retrieval mechanisms.

Potential retrieval channels:

```
Dense semantic retrieval
Sparse / lexical retrieval
Hybrid retrieval
Entity retrieval
Event retrieval
Graph retrieval
Hierarchical retrieval
Temporal retrieval
Metadata filtering
Source-aware retrieval
Evidence retrieval
Keyword retrieval
Chapter retrieval
Wiki-page retrieval
```

The system should dynamically choose which retrieval methods are appropriate for a query.

---

# 21. Adaptive Retrieval

Different questions require different retrieval strategies.

Example:

> "Who is Luffy?"

Simple entity retrieval may be sufficient.

But:

> "How did Luffy's relationship with Shanks evolve from Chapter 1 to the current story?"

requires:

```
entity retrieval
event retrieval
temporal retrieval
chapter retrieval
relationship traversal
synthesis
```

The agent should not perform maximum retrieval for every question.

Retrieval depth should adapt to query complexity.

---

# 22. Query Classification

The agent should classify questions into categories such as:

```
Factual
Entity
Relationship
Event
Temporal
Comparative
Multi-hop
Causal
Evidence
Theory
Canon
Contradiction
Summary
Recommendation
Spoiler-sensitive
Cross-series
```

Classification should determine retrieval strategy.

---

# 23. Query Decomposition

Complex queries should be decomposed.

Example:

> "Which characters have fought X, lost to X, later became allies with X, and were still alive after the final battle?"

Potential decomposition:

```
1. Identify X
2. Retrieve all combat events involving X
3. Filter outcomes
4. Resolve character entities
5. Retrieve later relationship events
6. Retrieve final battle
7. Apply temporal constraints
8. Apply status constraints
9. Verify evidence
10. Synthesize answer
```

---

# 24. Multi-Hop Reasoning

Manga Wiki should support reasoning across multiple entities and events.

Example:

> "Which characters knew both Roger and Luffy before Luffy entered the Grand Line?"

This may require:

```
Roger
  ↓
known_by
  ↓
candidate characters
  ↓
Luffy
  ↓
known_by
  ↓
intersection
  ↓
temporal filtering
  ↓
evidence verification
```

This cannot reliably be solved through a single top-k vector retrieval.

---

# 25. Agent Architecture — Product Requirements

The agent should have access to specialized tools.

Conceptual tools may include:

```
search_wiki()
search_entities()
search_chapters()
search_events()
search_relationships()
search_timeline()
traverse_graph()
search_sources()
retrieve_evidence()
verify_claim()
check_canon()
check_spoiler_boundary()
detect_conflicts()
```

The exact implementation is intentionally left to the technical architecture phase.

---

# 26. Agentic Retrieval

The agent should be able to:

1. Understand the question.
2. Determine what information is required.
3. Select retrieval strategies.
4. Execute searches.
5. Inspect intermediate results.
6. Identify missing evidence.
7. Perform additional retrieval.
8. Cross-check important claims.
9. Resolve contradictions.
10. Produce a final grounded answer.

The system should avoid unnecessary retrieval loops.

---

# 27. Evidence Verification

Before finalizing an answer, the system should ideally verify high-impact claims.

Example:

```
Claim:
"X defeated Y."
```

Verification:

```
Search event records
   ↓
Search chapter evidence
   ↓
Verify outcome
   ↓
Verify chronology
   ↓
Check conflicting sources
   ↓
Accept / qualify / reject
```

---

# 28. Citation Requirements

Citations should be:

* relevant
* precise
* traceable
* attached to claims
* preferably linked to the smallest useful evidence unit

A citation should not merely point to an enormous document when a chapter or section can be identified.

Ideal:

```
Claim
  [Chapter 123]
```

rather than:

```
Claim
  [Some 500-page source]
```

---

# 29. Spoiler-Aware Retrieval

Spoiler handling is a fundamental product capability.

The user may specify:

```
"I have read until Chapter 500."
```

The system should establish:

```
user_knowledge_boundary = Chapter 500
```

Retrieval should then exclude or appropriately mask evidence beyond that boundary.

The same concept should work for:

* volume
* chapter
* arc
* anime episode
* publication point

---

# 30. Spoiler Levels

Potential settings:

```
No spoiler beyond progress
Mild spoiler
Full spoiler
Custom chapter boundary
```

The system should never assume that a user wants spoilers.

---

# 31. Spoiler-Aware Answer Generation

Even if retrieval accidentally encounters later knowledge, the answer generator must not expose it when the user has established a boundary.

Therefore spoiler control should exist at multiple layers:

```
Retrieval filtering
    +
Context filtering
    +
Answer validation
```

---

# 32. Contradiction Detection

Manga information may conflict between:

* manga
* databooks
* author statements
* translations
* anime
* movies
* fan wikis

The system should detect conflicts rather than silently merging them.

Example:

```
Source A:
X was introduced in Chapter 100.

Source B:
X was first mentioned in Chapter 98.
```

The system should investigate whether:

* they refer to different events,
* one is incorrect,
* translation differences exist,
* the claim is ambiguous.

---

# 33. Source Authority

The retrieval system should be source-aware.

When multiple sources support a claim, the system should prefer stronger sources.

Conceptual ranking:

```
Primary canon
   >
Official supplementary material
   >
Author statements
   >
Secondary authoritative sources
   >
Community-maintained sources
   >
Speculation
```

The exact hierarchy must be configurable by manga.

---

# 34. Manga-Specific Knowledge Boundaries

Each manga may have different definitions of canon.

Therefore canon policy should not be hardcoded globally.

The system should support:

```
Manga-specific canon rules
Source-specific authority
Publication-specific rules
```

---

# 35. Wiki Page Generation

The system should be capable of generating or updating Wiki pages.

Potential page types:

```
Manga
Character
Chapter
Arc
Event
Location
Organization
Item
Ability
Timeline
Concept
```

Generated pages should contain:

* summary
* structured facts
* relationships
* timeline
* appearances
* evidence
* citations
* source metadata
* related pages

---

# 36. Human-in-the-Loop

AI-generated knowledge must remain inspectable.

Humans should eventually be able to:

* edit facts
* merge entities
* split entities
* correct relationships
* change source authority
* flag incorrect evidence
* approve generated pages
* revert changes

The system should maintain revision history.

---

# 37. Knowledge Revision

Knowledge should be versioned.

Conceptual model:

```
Version N
   ↓
New Source
   ↓
Re-evaluation
   ↓
Version N+1
```

The system should preserve historical changes where practical.

This is important because manga knowledge evolves as new chapters are published.

---

# 38. Incremental Knowledge Updating

The system should not rebuild an entire manga every time a new chapter arrives.

Preferred model:

```
Existing Knowledge
     +
New Chapter
     ↓
Incremental Extraction
     ↓
Entity Resolution
     ↓
Event Extraction
     ↓
Relationship Updates
     ↓
Timeline Update
     ↓
Evidence Update
     ↓
Retrieval Index Update
```

---

# 39. Persistent Knowledge

The system should accumulate knowledge over time.

A user asking the same question later should benefit from:

* previously compiled Wiki pages
* resolved entities
* known relationships
* previous source processing
* validated evidence

The goal is not merely persistent chat history.

It is **persistent domain knowledge**.

---

# 40. Cross-Manga Reasoning

This is a future capability.

The system may eventually support:

> "Compare the power systems of One Piece and Hunter × Hunter."

or:

> "Which manga protagonists have a relationship structure similar to X?"

However, cross-manga reasoning must not contaminate isolated manga knowledge.

Each manga should retain its own knowledge boundary.

---

# 41. Answer Quality Model

Manga Wiki should optimize for:

```
Accuracy
Evidence
Completeness
Relevance
Temporal correctness
Canon correctness
Spoiler safety
Citation correctness
Reasoning quality
```

Not merely:

```
Response fluency
```

---

# 42. Answer Types

The system should support different answer modes.

## Concise

For simple factual questions.

## Detailed

For normal research.

## Deep Research

For complex multi-hop questions.

## Evidence Mode

Prioritizes citations and source inspection.

## Timeline Mode

Organizes answer chronologically.

## Comparison Mode

Structured side-by-side analysis.

## Theory Mode

Separates:

```
Confirmed
Implied
Speculative
```

---

# 43. Example: Simple Query

Question:

> "Who is Shanks?"

Expected process:

```
Entity resolution
    ↓
Character retrieval
    ↓
Wiki page retrieval
    ↓
Evidence verification
    ↓
Answer
```

No unnecessary graph traversal should occur.

---

# 44. Example: Multi-Hop Query

Question:

> "Which characters have both interacted with Roger and later interacted with Luffy?"

Expected process:

```
Identify Roger
   ↓
Retrieve Roger relationships
   ↓
Identify candidate characters
   ↓
Identify Luffy
   ↓
Retrieve Luffy relationships
   ↓
Intersection
   ↓
Temporal filtering
   ↓
Evidence verification
   ↓
Answer
```

---

# 45. Example: Temporal Query

Question:

> "What events caused X to become Y?"

Expected process:

```
Identify X
   ↓
Identify Y
   ↓
Retrieve relevant events
   ↓
Sort by story chronology
   ↓
Detect causal relationships
   ↓
Verify evidence
   ↓
Construct causal narrative
```

The system must distinguish:

```
chronological sequence
```

from:

```
proven causality
```

It must not infer causality merely because event A occurred before event B.

---

# 46. Example: Canon Query

Question:

> "Is X actually canon?"

Expected answer structure:

```
Canon status:
Confirmed / Unconfirmed / Disputed

Primary evidence:
...

Secondary evidence:
...

Contradictions:
...

Conclusion:
...
```

---

# 47. Example: Theory Query

Question:

> "Is X secretly Y?"

The system should never answer simply:

> "Yes."

Instead:

```
Confirmed facts
    ↓
Evidence supporting theory
    ↓
Evidence against theory
    ↓
Alternative explanations
    ↓
Current canon status
    ↓
Assessment
```

The system should explicitly separate fact from interpretation.

---

# 48. Example: Spoiler Query

User:

> "I have read through Chapter 500. Explain why X did Y without spoilers."

The system should:

```
Detect boundary = 500
   ↓
Restrict retrieval
   ↓
Verify retrieved evidence
   ↓
Remove post-500 information
   ↓
Generate answer
```

---

# 49. Advanced RAG Research Program

The product should remain open to advanced retrieval techniques.

Potential candidates include:

```
Hybrid RAG
GraphRAG
Agentic RAG
Hierarchical retrieval
Entity-aware retrieval
Event-aware retrieval
Temporal retrieval
Query decomposition
Query rewriting
Multi-query retrieval
Reranking
MMR
Contextual retrieval
Knowledge compilation
Vectorless retrieval
PageIndex-style retrieval
Dynamic graph retrieval
Evidence retrieval
Adaptive retrieval
```

However:

> No technique should be adopted merely because it is considered "advanced."

Each technique must demonstrate measurable improvement against a benchmark.

---

# 50. Retrieval Research Principle

The project follows:

```
Hypothesis
   ↓
Implementation
   ↓
Benchmark
   ↓
Evaluation
   ↓
Adopt / Reject
```

rather than:

```
Popular technique
   ↓
Add to stack
```

This protects the project from unnecessary complexity.

---

# 51. Model Agnosticism

The system should avoid being permanently coupled to one LLM provider.

Potential model roles:

```
Fast model
Reasoning model
Extraction model
Embedding model
Reranker
Vision model
Classification model
```

The architecture should allow different models for different workloads.

---

# 52. Cost Awareness

Deep reasoning can be expensive.

The agent should use the cheapest sufficient strategy.

Conceptually:

```
Simple question
    ↓
Cheap retrieval + fast model

Complex question
    ↓
Multi-stage retrieval + stronger reasoning model
```

The system should avoid using maximum reasoning depth for trivial questions.

---

# 53. Latency Tiers

Potential target modes:

### Fast

Simple questions.

### Standard

Normal multi-hop questions.

### Deep

Complex research questions.

### Exhaustive

Maximum evidence retrieval and verification.

The user should eventually be able to choose depth.

---

# 54. Evaluation Framework

Evaluation is a first-class product requirement.

The project must maintain a Manga QA benchmark.

Initial benchmark categories:

```
Factual
Entity
Relationship
Multi-hop
Temporal
Causal
Comparative
Canon
Contradiction
Evidence
Spoiler
Theory
Long-context
```

---

# 55. Benchmark Dataset

An initial benchmark should contain at least:

```
100 factual questions
100 entity questions
100 relationship questions
100 multi-hop questions
100 temporal questions
100 evidence questions
100 canon questions
100 contradiction questions
100 spoiler-sensitive questions
```

The exact dataset size may change after experimentation.

---

# 56. Evaluation Metrics

Core metrics:

## Retrieval Recall

Did the system retrieve the relevant evidence?

## Retrieval Precision

How much retrieved information was actually useful?

## Answer Accuracy

Was the final answer correct?

## Citation Accuracy

Do citations actually support the claims?

## Citation Completeness

Are important claims supported?

## Faithfulness

Does the answer remain grounded in retrieved evidence?

## Temporal Accuracy

Are events ordered correctly?

## Canon Accuracy

Does the answer correctly classify canon status?

## Spoiler Safety

Did the answer reveal prohibited information?

## Entity Resolution Accuracy

Did the system identify the correct entities?

---

# 57. Hallucination Definition

For this project, hallucination is not merely "the model said something false."

A critical hallucination includes:

* unsupported factual claim
* fabricated chapter reference
* incorrect relationship
* incorrect event chronology
* false canon classification
* fabricated quotation
* conflation of entities
* inference presented as fact

These should be measured separately.

---

# 58. Quality Bar

A high-quality Manga Wiki answer should ideally satisfy:

```
Correct
Relevant
Grounded
Traceable
Canon-aware
Spoiler-safe
Temporally consistent
Appropriately qualified
```

Fluent but unsupported answers are not considered high quality.

---

# 59. Architecture Independence

This PRD intentionally does not prescribe:

* database technology
* graph database
* vector database
* embedding model
* LLM provider
* frontend framework
* backend framework
* orchestration framework
* deployment platform

These belong to the technical architecture phase.

---

# 60. Open-Source Foundation Strategy

The initial product strategy is:

> Start from an existing mature open-source knowledge/RAG platform rather than rebuilding generic infrastructure from scratch.

The primary candidate foundation is:

**WeKnora**

The product strategy is to investigate whether WeKnora can provide the base capabilities for:

* Wiki
* RAG
* Agent
* knowledge graph
* API
* MCP
* source management
* revision
* UI

before deciding what must be modified or replaced.

This is a **product hypothesis**, not an architectural commitment.

---

# 61. External Project Research

The technical research should compare ideas from:

### WeKnora

Primary foundation candidate.

Focus:

```
Wiki
Agent
RAG
Knowledge graph
MCP
Revision
Source management
```

### SAG

Research reference for:

```
entity retrieval
event retrieval
graph retrieval
relational reasoning
adaptive retrieval
```

### OpenKB

Research reference for:

```
knowledge compilation
persistent Wiki
structured knowledge
vectorless retrieval
navigable knowledge
```

### BYO-LLM-WIKI

Research reference for:

```
entity deduplication
synthesis pages
persistent knowledge
contextual retrieval
knowledge graph
```

### LLM-WIKI-RAG

Research reference for:

```
provenance
deterministic control
incremental updates
source lifecycle
```

### DeepAgents LLM Wiki

Research reference for:

```
agent-maintained Wiki
persistent knowledge
autonomous research
```

These projects should be treated as **sources of ideas**, not dependencies.

---

# 62. Technical Investigation Brief

A separate technical agent should inspect the candidate open-source repositories.

The technical investigation MUST answer:

## WeKnora

1. What is the actual architecture?
2. What components are reusable?
3. What components are tightly coupled?
4. How is Wiki implemented?
5. How is knowledge graph implemented?
6. How does retrieval work?
7. How does the agent work?
8. How are sources represented?
9. How is provenance represented?
10. How are revisions implemented?
11. How is MCP implemented?
12. How is the database structured?
13. How difficult is it to introduce Manga-specific entities?
14. How difficult is event modeling?
15. How difficult is temporal retrieval?
16. How difficult is spoiler filtering?
17. How difficult is source authority?
18. Where would a Manga ontology live?
19. What must be forked?
20. What should remain untouched?

## SAG

Determine:

* retrieval architecture
* graph architecture
* event model
* entity model
* query planning
* retrieval orchestration
* potential reusable ideas

## OpenKB

Determine:

* knowledge compilation pipeline
* Wiki representation
* retrieval approach
* vectorless retrieval
* indexing strategy

## BYO-LLM-WIKI

Determine:

* entity resolution
* entity deduplication
* synthesis
* memory
* graph construction

## LLM-WIKI-RAG

Determine:

* provenance
* deterministic update mechanisms
* incremental processing
* data lifecycle

---

# 63. Technical Agent Deliverable

The technical agent should produce a separate document:

```
docs/TECHNICAL-FEASIBILITY.md
```

It should contain:

```
Current Architecture
Reusable Components
Required Modifications
New Components
Technical Risks
Dependency Risks
Performance Risks
Data Model Proposal
Retrieval Proposal
Agent Proposal
Migration Strategy
Prototype Plan
Benchmark Plan
Estimated Complexity
```

It should NOT redefine the product vision.

---

# 64. Product vs Technical Boundary

The PO/Product layer decides:

```
What problem to solve
Why it matters
What capabilities are required
What quality means
What should be prioritized
```

The Technical layer decides:

```
How to implement it
Which technologies to use
Which repository to fork
Which architecture to adopt
Which algorithms perform best
```

Neither layer should silently replace the other.

---

# 65. MVP Definition

MVP should not attempt every advanced capability simultaneously.

A useful MVP should prove the central thesis:

> A domain-specific Wiki + structured knowledge + hybrid retrieval can answer manga questions more accurately and deeply than ordinary vector RAG.

Minimum viable system:

```
Manga ingestion
    ↓
Entity extraction
    ↓
Basic entity resolution
    ↓
Chapter knowledge
    ↓
Character knowledge
    ↓
Relationship knowledge
    ↓
Wiki pages
    ↓
Hybrid retrieval
    ↓
AI QA
    ↓
Evidence citations
```

---

# 66. MVP Should Include

### Knowledge

* Manga
* Character
* Chapter
* Arc
* Event
* Relationship
* Source
* Evidence

### Retrieval

* lexical
* semantic
* entity-aware
* chapter-aware

### Agent

* query classification
* basic decomposition
* retrieval orchestration
* evidence verification

### UI

* Wiki
* Search
* Chat
* Evidence

### Safety

* basic spoiler boundary
* source attribution

---

# 67. Post-MVP Capabilities

After MVP validation:

```
Knowledge graph traversal
Advanced multi-hop reasoning
Temporal reasoning
Canon classification
Contradiction detection
Advanced entity resolution
Knowledge compilation
Adaptive retrieval
Agentic research
Advanced spoiler system
Persistent user knowledge boundary
Cross-source synthesis
```

---

# 68. Long-Term Vision

The mature Manga Wiki should behave less like:

```
"Ask an LLM about manga"
```

and more like:

```
"Query an intelligent manga knowledge system."
```

A mature query could produce:

```
Answer
  ↓
Supporting evidence
  ↓
Relevant chapters
  ↓
Timeline
  ↓
Entities
  ↓
Relationships
  ↓
Contradictions
  ↓
Canon assessment
```

---

# 69. Example End-State Experience

User:

> "Explain every major change in the relationship between X and Y from their first meeting until the current chapter. Separate confirmed canon from interpretation, list the relevant chapters chronologically, and don't include information after chapter 900."

Manga Wiki internally:

```
Parse request
    ↓
Identify X and Y
    ↓
Boundary = Chapter 900
    ↓
Retrieve relationship events
    ↓
Retrieve relevant chapters
    ↓
Traverse relationship graph
    ↓
Sort by story chronology
    ↓
Detect relationship changes
    ↓
Retrieve evidence
    ↓
Classify canon status
    ↓
Detect contradictions
    ↓
Remove post-900 information
    ↓
Generate structured answer
    ↓
Verify citations
    ↓
Return answer
```

This represents the intended product direction.

---

# 70. Product Differentiation

Manga Wiki should not compete primarily on:

```
"We have a chatbot."
```

It should differentiate through:

### 1. Depth

Ability to answer multi-hop manga questions.

### 2. Evidence

Every important claim can be traced.

### 3. Structure

Knowledge is represented through entities, events, relationships and timelines.

### 4. Domain awareness

The system understands manga-specific concepts.

### 5. Temporal reasoning

The system understands story chronology.

### 6. Canon awareness

The system distinguishes authoritative knowledge from speculation.

### 7. Spoiler awareness

The system respects reader progress.

### 8. Persistent knowledge

The system compiles knowledge rather than repeatedly reconstructing it.

---

# 71. Product Principles

## Principle 1 — Evidence over confidence

A confident answer without evidence is inferior to a qualified answer with evidence.

## Principle 2 — Structure over chunks

Important knowledge should be represented structurally where possible.

## Principle 3 — Domain knowledge over generic RAG

Use manga-specific structure to improve retrieval and reasoning.

## Principle 4 — Benchmark before complexity

No advanced technique should be adopted without measurable benefit.

## Principle 5 — Facts and interpretations must remain separate

The system must clearly distinguish:

```
Fact
Inference
Interpretation
Theory
```

## Principle 6 — User spoiler boundary is sacred

Never sacrifice spoiler safety for answer completeness.

## Principle 7 — Human correction is part of the system

AI-generated knowledge must remain editable and auditable.

## Principle 8 — Persistent knowledge beats repeated reconstruction

The system should become smarter as its knowledge base grows.

---

# 72. Product Risks

## Risk 1 — Data quality

Bad source data produces bad knowledge.

Mitigation:

```
source authority
evidence
provenance
validation
```

## Risk 2 — Copyright / source licensing

The product must not rely on unauthorized redistribution of copyrighted manga material.

Mitigation:

* use appropriately licensed/authorized sources
* store metadata and derived knowledge where legally appropriate
* avoid unauthorized manga scan distribution
* design source connectors according to their terms

## Risk 3 — Hallucination

Mitigation:

```
evidence grounding
verification
citations
benchmark
```

## Risk 4 — Knowledge graph pollution

Mitigation:

```
entity resolution
deduplication
confidence
human review
```

## Risk 5 — Excessive complexity

Mitigation:

```
modular architecture
benchmark-driven adoption
phased roadmap
```

## Risk 6 — Cost

Mitigation:

```
adaptive models
caching
knowledge compilation
incremental updates
```

## Risk 7 — Latency

Mitigation:

```
retrieval tiers
precomputed knowledge
query classification
caching
```

---

# 73. Success Criteria

Manga Wiki is successful if users can reliably ask questions such as:

```
"Why?"
"How?"
"When?"
"Who?"
"What happened before?"
"What happened after?"
"Who knew whom?"
"What evidence supports this?"
"Is this canon?"
"What changed over time?"
"What are all the relevant chapters?"
"Does this theory have evidence?"
```

and receive answers that are:

```
accurate
grounded
traceable
spoiler-safe
structurally informative
```

---

# 74. North Star Metric

A potential North Star metric is:

> **Verified Deep Answer Rate**

Definition:

The percentage of complex benchmark questions for which the system produces an answer that is:

1. factually correct,
2. adequately supported by evidence,
3. correctly cited,
4. temporally consistent,
5. canon-aware where applicable,
6. spoiler-safe where applicable.

This is more meaningful than:

```
number of chats
number of Wiki pages
number of tokens
number of documents
```

---

# 75. Secondary Metrics

Potential metrics:

```
Retrieval Recall
Evidence Precision
Citation Accuracy
Answer Accuracy
Multi-hop Accuracy
Temporal Accuracy
Canon Accuracy
Spoiler Violation Rate
Entity Resolution Accuracy
Average Latency
Cost per Deep Query
Knowledge Update Latency
```

---

# 76. Initial Product Roadmap

## Phase 0 — Research

Objective:

Understand the available open-source foundations.

Deliverables:

```
WeKnora source analysis
SAG analysis
OpenKB analysis
BYO-WIKI analysis
Technical feasibility report
```

---

## Phase 1 — Foundation

Objective:

Create the Manga knowledge model on top of the selected foundation.

Deliverables:

```
Manga
Character
Chapter
Arc
Event
Relationship
Source
Evidence
```

---

## Phase 2 — Basic Manga RAG

Objective:

Prove that structured domain knowledge improves QA.

Deliverables:

```
ingestion
entity extraction
hybrid retrieval
Wiki
chat
citations
```

---

## Phase 3 — Deep Retrieval

Objective:

Enable multi-hop reasoning.

Deliverables:

```
graph retrieval
event retrieval
entity retrieval
query decomposition
reranking
evidence verification
```

---

## Phase 4 — Temporal + Canon

Objective:

Make the system genuinely manga-aware.

Deliverables:

```
timeline
temporal retrieval
canon model
source authority
contradiction detection
```

---

## Phase 5 — Spoiler Intelligence

Objective:

Make Manga Wiki safe for real readers.

Deliverables:

```
reader progress
chapter boundaries
spoiler-aware retrieval
spoiler-aware generation
spoiler validation
```

---

## Phase 6 — Autonomous Knowledge Engine

Objective:

Make knowledge continuously improve.

Deliverables:

```
incremental ingestion
knowledge compilation
Wiki auto-update
entity deduplication
knowledge validation
persistent research
```

---

# 77. Research Questions

The project should explicitly investigate:

1. Does a Manga-specific knowledge graph improve multi-hop QA?
2. Does event indexing improve temporal questions?
3. Does knowledge compilation outperform raw chunk retrieval?
4. Does Wiki navigation improve retrieval accuracy?
5. Does entity-aware retrieval reduce hallucination?
6. Does source authority improve factual correctness?
7. Does evidence verification reduce unsupported claims?
8. Does adaptive retrieval reduce cost without reducing accuracy?
9. Does vectorless retrieval help long manga documents?
10. Does GraphRAG outperform hybrid RAG on relationship questions?
11. Which reranking strategy performs best?
12. Which model is best for extraction versus reasoning?
13. How much knowledge should be precompiled versus retrieved dynamically?
14. What is the optimal granularity of evidence?
15. How should uncertain temporal information be represented?
16. How should canon status be modeled?
17. How can spoiler boundaries be enforced reliably?
18. How should conflicting translations be represented?
19. How should entity aliases be resolved across languages?
20. How should user corrections propagate through the knowledge graph?

---

# 78. Open Questions

The following are intentionally unresolved at PRD stage:

* Exact database architecture
* Exact graph technology
* Exact vector store
* Exact embedding model
* Exact reranker
* Exact LLM
* Exact agent framework
* Exact Wiki representation
* Exact source ingestion architecture
* Exact provenance schema
* Exact event schema
* Exact temporal representation
* Exact canon taxonomy
* Exact UI architecture

These should be resolved through technical investigation and experimentation.

---

# 79. Architectural Philosophy

The final system should ideally be:

```
Modular
Replaceable
Observable
Evaluatable
Source-grounded
Incrementally updatable
Model-agnostic
Domain-extensible
```

No single retrieval technique should become an irreversible dependency unless justified by evidence.

---

# 80. Decision Framework

When evaluating a new technology or technique, ask:

### Does it improve correctness?

If no → reject unless it provides another critical capability.

### Does it improve evidence grounding?

If yes → strong candidate.

### Does it improve deep reasoning?

If yes → benchmark.

### Does it increase complexity?

If yes → quantify the benefit.

### Does it create vendor lock-in?

If yes → evaluate alternatives.

### Can it be measured?

If no → avoid making it a core dependency.

---

# 81. Final Product Definition

Manga Wiki is:

> **An open, domain-specific knowledge engine that compiles manga information into a structured, evidence-grounded, navigable knowledge system and uses advanced retrieval and agentic reasoning to answer deep manga questions accurately, transparently, and with spoiler awareness.**

It is not simply:

```
Wiki + Chatbot
```

It is:

```
Knowledge Graph
    +
Compiled Wiki
    +
Evidence Store
    +
Multi-Strategy RAG
    +
Agentic Reasoning
    +
Temporal Model
    +
Canon Model
    +
Spoiler Model
    +
Human Curation
```

---

# 82. Immediate Next Step

The next action after this PRD is **not implementation**.

The next action is:

> Conduct a technical archaeology of WeKnora and compare it against the requirements defined here.

The technical agent should inspect the actual source code and determine:

```
What WeKnora already solves
    ↓
What can be extended
    ↓
What must be replaced
    ↓
What must be built
    ↓
What ideas should be borrowed from SAG
    ↓
What ideas should be borrowed from OpenKB
    ↓
What ideas should be borrowed from BYO-LLM-WIKI
    ↓
What architecture is realistically achievable
```

Only after this investigation should the project produce:

```
ARCHITECTURE.md
TECHNICAL-FEASIBILITY.md
DOMAIN-MODEL.md
RAG-ARCHITECTURE.md
EVALUATION.md
```

---

# 83. Product Constitution

The following statements are considered the core principles of Manga Wiki:

> **Manga Wiki is a knowledge engine, not merely a chatbot.**

> **Knowledge should be structured whenever structure improves reasoning.**

> **Every important factual claim should be traceable to evidence whenever possible.**

> **Facts, inference, interpretation, and speculation must not be silently conflated.**

> **Temporal reasoning is fundamental to manga knowledge.**

> **Canon status is part of knowledge, not merely metadata.**

> **Spoiler boundaries are a first-class constraint.**

> **The system should become more useful as its persistent knowledge improves.**

> **Advanced RAG techniques are hypotheses to be benchmarked, not decorations to be collected.**

> **The goal is not to build the most complicated RAG system.**

> **The goal is to build the most reliable system for answering deep questions about manga.**

---

# 84. End State

The ultimate ambition of Manga Wiki is that a user can ask:

> "I want to understand everything important about X, but I have only read up to Chapter 700."

and the system can function as:

```
Encyclopedia
↓
Research Assistant
↓
Timeline Engine
↓
Relationship Graph
↓
Evidence Database
↓
Canon Analyst
↓
Spoiler-Aware Guide
↓
Deep Reasoning Agent
```

while remaining grounded in the underlying evidence.

That is the product.

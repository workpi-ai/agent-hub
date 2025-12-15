---
description: Review code for bugs, security, and performance issues
---

Review the following code and identify critical issues:

## Code to Review
${*}

## Review Focus

1. **Bugs & Correctness**
   - Logic errors and edge cases
   - Null/undefined handling
   - Error handling completeness

2. **Security**
   - Input validation
   - Injection vulnerabilities
   - Sensitive data exposure

3. **Performance**
   - Algorithm complexity
   - Unnecessary computations
   - Memory leaks potential

4. **Code Quality Principles**
   - No magic numbers (use named constants)
   - DRY: No duplicated code or logic
   - SOLID: Single responsibility, proper abstractions
   - KISS: No over-engineering
   - YAGNI: No unnecessary features

5. **Project Consistency**
   - MUST use existing libraries in the project (no new dependencies without justification)
   - MUST follow existing code style and conventions
   - MUST follow similar implementation patterns already in the codebase
   - Check for similar implementations and ensure consistency

## Output Format

### Summary
One paragraph overview of code quality.

### Critical Issues
List issues that MUST be fixed before merging. Use severity:
- 🔴 **Critical**: Security vulnerabilities, data corruption risks
- 🟠 **High**: Bugs, performance bottlenecks

### Recommendations
Suggestions for improvement (Medium/Low severity).

### Positive Aspects
What the code does well.

Keep the review concise and actionable. Skip minor style issues unless they impact readability significantly.

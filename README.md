# Agent Hub

A centralized repository for AI agent configurations compatible with [GitHub Copilot Custom Agents](https://docs.github.com/en/copilot/reference/custom-agents-configuration) format and codev.

## Available Agents

| Agent | Description | File |
|-------|-------------|------|
| **Engineering** | Software engineering agent for coding tasks | [engineering.md](agents/engineering.md) |

## Agent File Format

Each agent is defined in a single Markdown file following the GitHub Copilot Custom Agents specification:

```markdown
---
name: Agent Name
description: Agent description and capabilities
tools: ["List", "Read", "Edit", "Write", "Grep", "Glob", "Shell", "TodoWrite", "WebSearch", "WebFetch", "Diagnostics"]
---

Agent instructions and system prompt content...
```

### Properties

- **name** (string): Display name for the custom agent
- **description** (string, required): Description of the agent's purpose and capabilities
- **tools** (list of strings): List of tool names the agent can use
  - Use `["*"]` to enable all available tools
  - Use specific tool names for selective access

## Available Tools

| Tool Name | Purpose |
|-----------|---------|
| `List` | List directory contents |
| `Read` | Read file contents |
| `Edit` | Edit existing files |
| `Write` | Create new files |
| `Grep` | Search text in files |
| `Glob` | Find files by pattern |
| `Shell` | Execute shell commands |
| `TodoWrite` | Manage task lists |
| `WebSearch` | Search the web |
| `WebFetch` | Fetch web content |
| `Diagnostics` | Get LSP diagnostics |

## Usage

### With codev

```bash
# Clone agent-hub to codev's registry directory
mkdir -p ~/.codev/agents/registry
git clone https://github.com/workpi-ai/agent-hub ~/.codev/agents/registry/agent-hub

# List available agents
codev agent list

# Use an agent
codev agent use engineering
```

## Creating Custom Agents

Create a `.md` file with YAML frontmatter and agent instructions:

```markdown
---
name: My Custom Agent
description: Specialized agent for my specific use case
tools: ["Read", "Edit", "Shell"]
---

You are a specialized agent for...

## Guidelines

- Guideline 1
- Guideline 2
```

## References

- [GitHub Copilot Custom Agents Documentation](https://docs.github.com/en/copilot/reference/custom-agents-configuration)
- [codev](https://github.com/workpi-ai/codev) - AI-powered coding CLI

## License

See [LICENSE](LICENSE) for details.

First off, thanks for taking the time to contribute! 💪🐘🎉

The following is a set of guidelines for contributing to Database Lab Engine (DLE) and its additional components, which are hosted on GitLab and GitHub:
- https://gitlab.com/postgres-ai/database-lab
- https://github.com/postgres-ai/database-lab-engine

These are mostly guidelines, not rules. Use your best judgment, and feel free to propose changes to this document in a pull request.

---

#### Table Of Contents

[Code of Conduct](#code-of-conduct)

[TL;DR – I just have a question, where to ask it?](#tldr-i-just-have-a-question-where-to-ask-it)

[How Can I Contribute?](#how-can-i-contribute)
  * [Reporting Bugs](#reporting-bugs)
  * [Proposing Enhancements](#proposing-enhancements)
  * [Your First Code Contribution](#your-first-code-contribution)
  * [Merge Requests / Pull Requests](#xxx)

[Repo Overview](#repo-overview)

[Styleguides](#styleguides)
  * [Git Commit Messages](#git-commit-messages)
  * [Go Styleguide](#go-styleguide)
  * [Documentation Styleguide](#documentation-styleguide)

---

## Code of Conduct
This project and everyone participating in it is governed by the [Database Lab Engine Community Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code. Please report unacceptable behavior to [community@postgres.ai](mailto:community@postgres.ai).

## TL;DR – I just have a question, where to ask it?

> **Note:** Please don't open an issue to just ask a question. You will get faster results by using the channels below.

- Fast ways to get in touch: [Contact page](https://postgres.ai/contact)
- [Database Lab Community Slack](https://slack.postgres.ai) (English)
- [Telegram](https://t.me/databaselabru) (Russian)

## How Can I Contribute?
### Reporting Bugs
- Use a clear and descriptive title for the issue to identify the problem.
- Make sure you test against the latest released version. It is possible that we may have already fixed the bug you're experiencing.
- Provide steps to reproduce the issue, including DLE version, PostgreSQL version, and the platform you are running on (some examples: RDS, self-managed Postgres on an EC2 instance, self-managed Postgres on-prem).
- Explain which behavior you expected to see instead and why.
- Please include DLE logs. Include Postgres logs from clones and/or the sync container's Postgres, if relevant.
- Describe DLE configuration: mode you are using (physical or logical), other details of DLE configuration, Postgres configuration snippets.
- If the issue is related to UI, include screenshots and animated GIFs. Please, do NOT use screenshots for console output, configs, and logs – for those, always prefer the textual form.
- Check if you have some sensitive information in the logs and configs and remove, if any.
- You can submit a bug report in either [GitLab Issues](https://gitlab.com/postgres-ai/database-lab) or [GitHub Issues](https://github.com/postgres-ai/database-lab-engine) sections – both places are monitored.
- If you believe that there is an urgency related to the reported bug, feel free to reach out to the project maintainers additionally, using one of [the channels described above](#tldr-i-just-have-a-question-where-to-ask-it).

### Proposing Enhancements
This section guides you through submitting an enhancement suggestion for DLE, including completely new features and minor improvements to existing functionality. Following these guidelines helps maintainers and the community understand your suggestion and find related proposals.

When you are creating an enhancement suggestion, please include as many details as possible. Include the steps that you imagine you would take if the feature you're requesting existed.

Enhancement suggestions are tracked on [GitLab](https://gitlab.com/postgres-ai/database-lab) or [GitHub](https://github.com/postgres-ai/database-lab-engine). Recommendations:

- Use a clear and descriptive title for the issue to identify the suggestion.
- Provide a step-by-step description of the proposed enhancement in as many details as possible.
- Provide specific examples to demonstrate the steps. Include copy/pasteable snippets which you use in those examples
- Use Markdown to format code snippets and improve the overall look of your ussues (Markdown docs: [GitLab](https://docs.gitlab.com/ee/user/markdown.html), [GitHub](https://github.github.com/gfm/)).
- Describe the current behavior and explain which behavior you expected to see instead and why.
- If your proposal is related to UI, include screenshots and animated GIFs which help you demonstrate the steps or point out the part of DLE which the suggestion is related to. Please, do NOT use screenshots for console output, logs, configs.
- Explain why this proposal would be useful to most DLE users.
- Specify which version of DLE you're using. If it makes sense, specify Postgres versions too.
- Specify the name and version of the OS you're using.

### Your First Code Contribution
We appreciate first time contributors and we are happy to assist you in getting started. In case of questions, just reach out to us!

You find some issues that are considered as good for first time contributors looking at [the issues with the `good-first-issue` label](https://gitlab.com/postgres-ai/database-lab/-/issues?label_name%5B%5D=good+first+issue).
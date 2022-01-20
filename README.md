<p><img width="500" src="./assets/dle.svg" border="0" /></p>

---

# Database Lab Engine (DLE)
Database Lab Engine (DLE) is an open source technology that allows blazing-fast cloning of Postgres databases:

- Build dev/QA/staging environments based on full-size production-like databases.
- Provide temporary full-size database clones for SQL query analysis and optimization (see also: [SQL optimization chatbot Joe](https://gitlab.com/postgres-ai/joe)).
- Automatically test database changes in CI/CD pipelines to avoid incidents in production.

For example, cloning a 1 TiB PostgreSQL database can take less than 2 seconds and dozens of independent clones can run on a single machine without extra costs, supporting lots of development and testing activities.

<p><img src="./assets/dle-demo-animated.gif" border="0" /></p>

## How it works
Thin cloning is fast because it is based on the [CoW (Copy-on-Write)](https://en.wikipedia.org/wiki/Copy-on-write#In_computer_storage). DLE supports two technologies to enable CoW and thin cloning: [ZFS](https://en.wikipedia.org/wiki/ZFS) (default) and [LVM](https://en.wikipedia.org/wiki/Logical_Volume_Manager_(Linux)).

With ZFS, Database Lab Engine periodically creates a new snapshot of the data directory, and maintains a set of snapshots, periodically deleting the old ones. When requesting a new clone, users choose which snapshot to use. For example, one can request to create two clones, one with a very fresh database state, and another corresponding to yesterday morning. In a few seconds, clones are created, and it immediately becomes possible to compare two database versions: data, SQL query plans, and so on.

Read more:
- [How it works](https://postgres.ai/products/how-it-works)
- [Database Migration Testing](https://postgres.ai/products/database-migration-testing)
- [SQL Optimization with Joe Bot](https://postgres.ai/products/joe)
- [Questions and answers](https://postgres.ai/docs/questions-and-answers)

## Where to start
- [Database Lab tutorial for any PostgreSQL database](https://postgres.ai/docs/tutorials/database-lab-tutorial)
- [Database Lab tutorial for Amazon RDS](https://postgres.ai/docs/tutorials/database-lab-tutorial-amazon-rds)
- [Terraform module template (AWS)](https://postgres.ai/docs/how-to-guides/administration/install-database-lab-with-terraform)

## Case studies
- Qiwi: [How Qiwi Controls the Data to Accelerate Development](https://postgres.ai/resources/case-studies/qiwi)
- GitLab: [How GitLab iterates on SQL performance optimization workflow to reduce downtime risks](https://postgres.ai/resources/case-studies/gitlab)

## Features
...

## How to contribute
> Please support the project giving a GitLab star! It's on [the main page](https://gitlab.com/postgres-ai/database-lab), at the upper right corner:
>
> ![Add a star](./assets/star.gif)

Lighter backgrounds:
<p>
  <img width="200px" src="https://gitlab.com/postgres-ai/database-lab/-/raw/rework-readme/assets/powered-by-dle-for-light-background.svg" />
</p>

```html
<a href="http://databaselab.io">
  <img width="200px" src="https://gitlab.com/postgres-ai/database-lab/-/raw/rework-readme/assets/powered-by-dle-for-light-background.svg" />
</a>
```

Darker backgrounds:
<p style="background-color: #bbb">
  <img width="200px" src="https://gitlab.com/postgres-ai/database-lab/-/raw/rework-readme/assets/powered-by-dle-for-dark-background.svg" />
</p>

```html
<a href="http://databaselab.io">
  <img width="200px" src="https://gitlab.com/postgres-ai/database-lab/-/raw/rework-readme/assets/powered-by-dle-for-dark-background.sgv" />
</a>
```

Open [an Issue](https://gitlab.com/postgres-ai/database-lab/-/issues) to discuss ideas, open [a Merge Request](https://gitlab.com/postgres-ai/database-lab/-/merge_requests) to propose a change.

... tbd ...
See our [GitLab Container Registry](https://gitlab.com/postgres-ai/database-lab/container_registry) to find the images built for development branches.
<!-- TODO: SDK docs -->
<!-- TODO: Contribution guideline -->

Development requirements:
1. Install `golangci-lint`: https://github.com/golangci/golangci-lint#install

### Reference guides
- [DLE components](https://postgres.ai/docs/reference-guides/database-lab-engine-components)
- [DLE configuration reference](https://postgres.ai/docs/database-lab/config-reference)
- [DLE API reference](https://postgres.ai/swagger-ui/dblab/)
- [Client CLI reference](https://postgres.ai/docs/database-lab/cli-reference)

### How-to guides
- [How to install Database Lab with Terraform on AWS](https://postgres.ai/docs/how-to-guides/administration/install-database-lab-with-terraform)
- [How to install and initialize Database Lab CLI](https://postgres.ai/docs/guides/cli/cli-install-init)
- [How to manage DLE](https://postgres.ai/docs/how-to-guides/administration)
- [How to work with clones](https://postgres.ai/docs/how-to-guides/cloning)

### Miscellaneous
- [DLE Docker images](https://hub.docker.com/r/postgresai/dblab-server)
- [Extended Docker images for PostgreSQL (with plenty of extensions)](https://hub.docker.com/r/postgresai/extended-postgres)
- [SQL Optimization chatbot (Joe Bot)](https://postgres.ai/docs/joe-bot)
- [DB Migration Checker](https://postgres.ai/docs/db-migration-checker)

## License
DLE source code is licensed under the OSI-approved open source license GNU Affero General Public License version 3 (AGPLv3).

Reach out to the Postgres.ai team if you want a trial or commercial license that does not contain the GPL clauses: [Contact page](https://postgres.ai/contact).

## Where to get help
- [Contact page](https://postgres.ai/contact)
- [Community Slack](https://slack.postgres.ai)

<!-- 
## Translations
- ...
-->

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
### Give the project a star
The easiest way to contribute is to give the project a GitHub/GitLab star:

![Add a star](./assets/star.gif)

### Mention that you use DLE
Please post a tweet mentioning [@Database_Lab](https://twitter.com/Database_Lab) or share the link to this repo in your favorite social network.

If you are actively using DLE at work, think where could you mention it. The best way of mentioning is using graphics with a link. Brand assets can be found in the `./assets` folder. Feel free to put them to your documents, slide decks, application and website interfaces to show that you use DLE.

HTML snippet for lighter backgrounds:
<p>
  <img width="400" src="https://gitlab.com/postgres-ai/database-lab/-/raw/rework-readme/assets/powered-by-dle-for-light-background.svg" />
</p>

```html
<a href="http://databaselab.io">
  <img width="400" src="https://gitlab.com/postgres-ai/database-lab/-/raw/rework-readme/assets/powered-by-dle-for-light-background.svg" />
</a>
```

Fro darker backgrounds:
<p style="background-color: #bbb">
  <img width="400" src="https://gitlab.com/postgres-ai/database-lab/-/raw/rework-readme/assets/powered-by-dle-for-dark-background.svg" />
</p>

```html
<a href="http://databaselab.io">
  <img width="400" src="https://gitlab.com/postgres-ai/database-lab/-/raw/rework-readme/assets/powered-by-dle-for-dark-background.sgv" />
</a>
```

### Propose an idea or report a bug
Check out our [contributing guide](./CONTRIBUTING.md) for more details.

### Participate in development
Check out our [contributing guide](./CONTRIBUTING.md) for more details.

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

More you can found in [the "How-to guides" section](https://postgres.ai/docs/how-to-guides) of the docs. 

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

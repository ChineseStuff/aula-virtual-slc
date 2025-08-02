# Neo4j Graph Database for Traditional Chinese Medicine Herbs

This folder contains data and initialization scripts for the Neo4j graph database used in this project.

## Overview

The Neo4j database models herbs and their relationships within the context of Traditional Chinese Medicine (TCM). The graph structure allows for flexible representation of complex connections, such as:

- **Herbs**: Nodes representing individual medicinal herbs.
- **Properties**: Nodes for properties like flavor, nature, and meridian tropism.
- **Formulas**: Nodes for classical TCM formulas containing multiple herbs.
- **Relationships**: Edges describing how herbs relate to each other, their properties, and their roles in formulas (e.g., "used_in", "has_property", "paired_with").

## Usage

- The database is initialized using the files in the [`init`](neo4j/init/) directory.
- Data files are stored in the [`data`](neo4j/data/) directory.

## Example Relationships

- Herb A `has_property` Bitter
- Herb B `used_in` Formula X
- Herb A `paired_with` Herb B

This structure enables advanced queries about herb compatibility, formula composition, and TCM theory.

---
For more details, see the main project [README.md](../README.md).
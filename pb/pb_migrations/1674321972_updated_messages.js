migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("utxo8gdx4ndqsqz")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "eqs7ijbx",
    "name": "job_title",
    "type": "text",
    "required": false,
    "unique": false,
    "options": {
      "min": null,
      "max": null,
      "pattern": ""
    }
  }))

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "gsujvpsq",
    "name": "job_role",
    "type": "text",
    "required": false,
    "unique": false,
    "options": {
      "min": null,
      "max": null,
      "pattern": ""
    }
  }))

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "1kuysu5p",
    "name": "business_type",
    "type": "text",
    "required": false,
    "unique": false,
    "options": {
      "min": null,
      "max": null,
      "pattern": ""
    }
  }))

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "oouivgfj",
    "name": "company_name",
    "type": "text",
    "required": false,
    "unique": false,
    "options": {
      "min": null,
      "max": null,
      "pattern": ""
    }
  }))

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "wb03lrqa",
    "name": "company_type",
    "type": "text",
    "required": false,
    "unique": false,
    "options": {
      "min": null,
      "max": null,
      "pattern": ""
    }
  }))

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "lgmfpdka",
    "name": "company_size",
    "type": "text",
    "required": false,
    "unique": false,
    "options": {
      "min": null,
      "max": null,
      "pattern": ""
    }
  }))

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "wclinpqr",
    "name": "industry",
    "type": "text",
    "required": false,
    "unique": false,
    "options": {
      "min": null,
      "max": null,
      "pattern": ""
    }
  }))

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "dyswgfo4",
    "name": "country",
    "type": "text",
    "required": false,
    "unique": false,
    "options": {
      "min": null,
      "max": null,
      "pattern": ""
    }
  }))

  return dao.saveCollection(collection)
}, (db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("utxo8gdx4ndqsqz")

  // remove
  collection.schema.removeField("eqs7ijbx")

  // remove
  collection.schema.removeField("gsujvpsq")

  // remove
  collection.schema.removeField("1kuysu5p")

  // remove
  collection.schema.removeField("oouivgfj")

  // remove
  collection.schema.removeField("wb03lrqa")

  // remove
  collection.schema.removeField("lgmfpdka")

  // remove
  collection.schema.removeField("wclinpqr")

  // remove
  collection.schema.removeField("dyswgfo4")

  return dao.saveCollection(collection)
})

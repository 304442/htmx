migrate((db) => {
  const dao = new Dao(db)
  const collection = dao.findCollectionByNameOrId("utxo8gdx4ndqsqz")

  // add
  collection.schema.addField(new SchemaField({
    "system": false,
    "id": "lmidkioi",
    "name": "use_case",
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
  collection.schema.removeField("lmidkioi")

  return dao.saveCollection(collection)
})

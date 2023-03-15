package migrations

import (
	"encoding/json"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		jsonData := `{
			"id": "5wv76wa1hbpu6bl",
			"created": "2023-03-15 20:58:44.827Z",
			"updated": "2023-03-15 20:58:44.827Z",
			"name": "orders",
			"type": "base",
			"system": false,
			"schema": [
				{
					"system": false,
					"id": "io2namdi",
					"name": "nftId",
					"type": "number",
					"required": false,
					"unique": false,
					"options": {
						"min": null,
						"max": null
					}
				},
				{
					"system": false,
					"id": "crrp9ykl",
					"name": "user",
					"type": "relation",
					"required": true,
					"unique": false,
					"options": {
						"collectionId": "_pb_users_auth_",
						"cascadeDelete": false,
						"minSelect": null,
						"maxSelect": 1,
						"displayFields": []
					}
				},
				{
					"system": false,
					"id": "nswggmd2",
					"name": "paid",
					"type": "bool",
					"required": false,
					"unique": false,
					"options": {}
				}
			],
			"listRule": null,
			"viewRule": null,
			"createRule": null,
			"updateRule": null,
			"deleteRule": null,
			"options": {}
		}`

		collection := &models.Collection{}
		if err := json.Unmarshal([]byte(jsonData), &collection); err != nil {
			return err
		}

		return daos.New(db).SaveCollection(collection)
	}, func(db dbx.Builder) error {
		dao := daos.New(db);

		collection, err := dao.FindCollectionByNameOrId("5wv76wa1hbpu6bl")
		if err != nil {
			return err
		}

		return dao.DeleteCollection(collection)
	})
}

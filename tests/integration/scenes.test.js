const grpc = require('grpc');

const organizations_pb = require('./protos/organizations/organizations_pb');
const organizations_grpc_pb = require('./protos/organizations/organizations_grpc_pb');

const scenes_pb = require('./protos/scenes/scenes_pb');
const scenes_grpc_pb = require('./protos/scenes/scenes_grpc_pb');

let SERVER = process.env.SERVER_HOST + ':' + process.env.SERVER_PORT;

let checkDatastore = () => {
    return new Promise((resolve, reject) => {
        let datastore = require('@google-cloud/datastore')({
            projectId: process.env.GOOGLE_PROJECT_ID,
            keyFilename: process.env.GOOGLE_APPLICATION_CREDENTIALS,
        });

        let query = datastore.createQuery('TestCheck');

        datastore.runQuery(query, (err, entities) => {
            if (err) {
                reject();
            } else {
                resolve(datastore);
            }
        });
    });
};

describe('Scenes', () => {
    let ORG_ID = null;

    beforeAll(() => {
        return checkDatastore().then(() => {
            let client = new organizations_grpc_pb.OrganizationsClient(SERVER, grpc.credentials.createInsecure());

            let createRequest = new organizations_pb.CreateOrganizationRequest();
            let org = new organizations_pb.Organization();
            org.setName("Test Org.");
            createRequest.setOrganization(org);

            return new Promise((resolve, reject) => {
                client.create(createRequest, (err, response) => {
                    expect(err).toEqual(null);
                    expect(response.getName()).toEqual("Test Org.");

                    ORG_ID = response.getId();

                    resolve()
                })
            });
        });
    });


    test('Create', () => {
        let client = new scenes_grpc_pb.ScenesClient(SERVER, grpc.credentials.createInsecure());

        let createRequest = new scenes_pb.CreateSceneRequest();
        createRequest.setOrganizationId(ORG_ID);

        let scene = new scenes_pb.Scene();
        scene.setName("Test Co.");
        createRequest.setScene(scene);

        return new Promise((resolve, reject) => {
            client.create(createRequest, (err, response) => {
                expect(err).toEqual(null);
                expect(response.getName()).toEqual("Test Co.");
                resolve()
            })
        });
    });

    test('Get', () => {
        let client = new scenes_grpc_pb.ScenesClient(SERVER, grpc.credentials.createInsecure());

        let createRequest = new scenes_pb.CreateSceneRequest();
        createRequest.setOrganizationId(ORG_ID);

        let scene = new scenes_pb.Scene();
        scene.setName("Test Co.");
        createRequest.setScene(scene);

        return new Promise((resolve, reject) => {
            client.create(createRequest, (err, response) => {
                expect(err).toEqual(null);
                expect(response.getName()).toEqual("Test Co.");
                resolve(response.getId())
            })
        }).then((sceneId) => {
            let getRequest = new scenes_pb.GetSceneRequest();
            getRequest.setOrganizationId(ORG_ID);
            getRequest.setSceneId(sceneId);

            return new Promise((resolve, reject) => {
                client.get(getRequest, (err, response) => {
                    expect(err).toEqual(null);
                    expect(response.getName()).toEqual("Test Co.");
                    resolve()
                })
            });
        })
    });

    test('List', () => {
        let client = new scenes_grpc_pb.ScenesClient(SERVER, grpc.credentials.createInsecure());

        let listRequest = new scenes_pb.ListScenesRequest();
        listRequest.setOrganizationId(ORG_ID);

        return new Promise((resolve, reject) => {
            client.list(listRequest, (err, response) => {
                expect(err).toEqual(null);
                expect(response.getScenesList().length).toBeGreaterThan(0);
                resolve()
            })
        });
    });

    test('Update', () => {
        let client = new scenes_grpc_pb.ScenesClient(SERVER, grpc.credentials.createInsecure());

        let createRequest = new scenes_pb.CreateSceneRequest();
        createRequest.setOrganizationId(ORG_ID);

        let scene = new scenes_pb.Scene();
        scene.setName("Test Co.");
        createRequest.setScene(scene);

        return new Promise((resolve, reject) => {
            client.create(createRequest, (err, response) => {
                expect(err).toEqual(null);
                expect(response.getName()).toEqual("Test Co.");
                resolve(response.getId())
            })
        }).then((sceneId) => {
            let updateRequest = new scenes_pb.UpdateSceneRequest();
            updateRequest.setOrganizationId(ORG_ID);
            updateRequest.setSceneId(sceneId);

            let updateScene = new scenes_pb.Scene();
            updateScene.setName("New Name");
            updateRequest.setScene(updateScene);

            return new Promise((resolve, reject) => {
                client.update(updateRequest, (err, response) => {
                    expect(err).toEqual(null);
                    expect(response.getName()).toEqual("New Name");
                    resolve()
                })
            });
        });
    });

    test('Delete', () => {
        let client = new scenes_grpc_pb.ScenesClient(SERVER, grpc.credentials.createInsecure());

        let createRequest = new scenes_pb.CreateSceneRequest();
        createRequest.setOrganizationId(ORG_ID);

        let scene = new scenes_pb.Scene();
        scene.setName("Test Co.");
        createRequest.setScene(scene);

        return new Promise((resolve, reject) => {
            client.create(createRequest, (err, response) => {
                expect(err).toEqual(null);
                expect(response.getName()).toEqual("Test Co.");
                resolve(response.getId())
            })
        }).then((sceneId) => {
            let deleteRequest = new scenes_pb.DeleteSceneRequest();
            deleteRequest.setOrganizationId(ORG_ID);
            deleteRequest.setSceneId(sceneId);

            return new Promise((resolve, reject) => {
                client.delete(deleteRequest, (err, response) => {
                    expect(err).toEqual(null);
                    expect(response.toArray().length).toBe(0);
                    resolve()
                })
            });
        })
    });
});
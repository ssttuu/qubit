const grpc = require('grpc');

const organizations_pb = require('./protos/organizations/organizations_pb');
const organizations_grpc_pb = require('./protos/organizations/organizations_grpc_pb');

const scenes_pb = require('./protos/scenes/scenes_pb');
const scenes_grpc_pb = require('./protos/scenes/scenes_grpc_pb');

const operators_pb = require('./protos/operators/operators_pb');
const operators_grpc_pb = require('./protos/operators/operators_grpc_pb');

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

describe('Operators', () => {
    let ORG_ID = null;
    let SCENE_ID = null;

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

                    resolve(ORG_ID)
                })
            });
        }).then(() => {
            let client = new scenes_grpc_pb.ScenesClient(SERVER, grpc.credentials.createInsecure());

            let createRequest = new scenes_pb.CreateSceneRequest();
            createRequest.setOrganizationId(ORG_ID);
            let scene = new scenes_pb.Scene();
            scene.setName("Test Scene");
            createRequest.setScene(scene);

            return new Promise((resolve, reject) => {
                client.create(createRequest, (err, response) => {
                    expect(err).toEqual(null);
                    expect(response.getName()).toEqual("Test Scene");

                    SCENE_ID = response.getId();

                    resolve(SCENE_ID)
                })
            });
        });
    });


    test('Create', () => {
        let client = new operators_grpc_pb.OperatorsClient(SERVER, grpc.credentials.createInsecure());

        let createRequest = new operators_pb.CreateOperatorRequest();
        createRequest.setOrganizationId(ORG_ID);
        createRequest.setSceneId(SCENE_ID);

        let operator = new operators_pb.Operator();
        operator.setName("Test Co.");
        createRequest.setOperator(operator);

        return new Promise((resolve, reject) => {
            client.create(createRequest, (err, response) => {
                expect(err).toEqual(null);
                expect(response.getName()).toEqual("Test Co.");
                resolve()
            })
        });
    });

    test('Get', () => {
        let client = new operators_grpc_pb.OperatorsClient(SERVER, grpc.credentials.createInsecure());

        let createRequest = new operators_pb.CreateOperatorRequest();
        createRequest.setOrganizationId(ORG_ID);
        createRequest.setSceneId(SCENE_ID);

        let operator = new operators_pb.Operator();
        operator.setName("Test Co.");
        createRequest.setOperator(operator);

        return new Promise((resolve, reject) => {
            client.create(createRequest, (err, response) => {
                expect(err).toEqual(null);
                expect(response.getName()).toEqual("Test Co.");
                resolve(response.getId())
            })
        }).then((operatorId) => {
            let getRequest = new operators_pb.GetOperatorRequest();
            getRequest.setOrganizationId(ORG_ID);
            getRequest.setSceneId(SCENE_ID);
            getRequest.setOperatorId(operatorId);

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
        let client = new operators_grpc_pb.OperatorsClient(SERVER, grpc.credentials.createInsecure());

        let listRequest = new operators_pb.ListOperatorsRequest();
        listRequest.setOrganizationId(ORG_ID);
        listRequest.setSceneId(SCENE_ID);

        return new Promise((resolve, reject) => {
            client.list(listRequest, (err, response) => {
                expect(err).toEqual(null);
                expect(response.getOperatorsList().length).toBeGreaterThan(0);
                resolve()
            })
        });
    });

    test('Update', () => {
        let client = new operators_grpc_pb.OperatorsClient(SERVER, grpc.credentials.createInsecure());

        let createRequest = new operators_pb.CreateOperatorRequest();
        createRequest.setOrganizationId(ORG_ID);
        createRequest.setSceneId(SCENE_ID);

        let operator = new operators_pb.Operator();
        operator.setName("Test Co.");
        createRequest.setOperator(operator);

        return new Promise((resolve, reject) => {
            client.create(createRequest, (err, response) => {
                expect(err).toEqual(null);
                expect(response.getName()).toEqual("Test Co.");
                resolve(response.getId())
            })
        }).then((operatorId) => {
            let updateRequest = new operators_pb.UpdateOperatorRequest();
            updateRequest.setOrganizationId(ORG_ID);
            updateRequest.setSceneId(SCENE_ID);
            updateRequest.setOperatorId(operatorId);

            let updateOperator = new operators_pb.Operator();
            updateOperator.setName("New Name");
            updateRequest.setOperator(updateOperator);

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
        let client = new operators_grpc_pb.OperatorsClient(SERVER, grpc.credentials.createInsecure());

        let createRequest = new operators_pb.CreateOperatorRequest();
        createRequest.setOrganizationId(ORG_ID);
        createRequest.setSceneId(SCENE_ID);

        let operator = new operators_pb.Operator();
        operator.setName("Test Co.");
        createRequest.setOperator(operator);

        return new Promise((resolve, reject) => {
            client.create(createRequest, (err, response) => {
                expect(err).toEqual(null);
                expect(response.getName()).toEqual("Test Co.");
                resolve(response.getId())
            })
        }).then((operatorId) => {
            let deleteRequest = new operators_pb.DeleteOperatorRequest();
            deleteRequest.setOrganizationId(ORG_ID);
            deleteRequest.setSceneId(SCENE_ID);
            deleteRequest.setOperatorId(operatorId);

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
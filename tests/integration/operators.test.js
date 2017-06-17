const grpc = require('grpc');

const helpers = require('./lib/helpers');
const operators_pb = require('./protos/operators/operators_pb');
const operators_grpc_pb = require('./protos/operators/operators_grpc_pb');
const organizations_pb = require('./protos/organizations/organizations_pb');
const organizations_grpc_pb = require('./protos/organizations/organizations_grpc_pb');
const projects_pb = require('./protos/projects/projects_pb');
const projects_grpc_pb = require('./protos/projects/projects_grpc_pb');
const scenes_pb = require('./protos/scenes/scenes_pb');
const scenes_grpc_pb = require('./protos/scenes/scenes_grpc_pb');
//const geometry_pb = require('./protos/geometry/geometry_pb');

const SERVER = process.env.API_SERVICE_ADDRESS;

describe('Operators', () => {
    let ORGANIZATIONS_CLIENT = null;
    let PROJECTS_CLIENT = null;
    let SCENES_CLIENT = null;
    let OPERATORS_CLIENT = null;
    let ORG_ID = null;
    let PROJECT_ID = null;
    let SCENE_ID = null;

    beforeAll(() => {
        ORGANIZATIONS_CLIENT = new organizations_grpc_pb.OrganizationsClient(SERVER, grpc.credentials.createInsecure());
        PROJECTS_CLIENT = new projects_grpc_pb.ProjectsClient(SERVER, grpc.credentials.createInsecure());
        SCENES_CLIENT = new scenes_grpc_pb.ScenesClient(SERVER, grpc.credentials.createInsecure());
        OPERATORS_CLIENT = new operators_grpc_pb.OperatorsClient(SERVER, grpc.credentials.createInsecure());
        return helpers.createOrganization(ORGANIZATIONS_CLIENT, 'Test Organization').then((organization) => {
            ORG_ID = organization.getId();
            return helpers.createProject(PROJECTS_CLIENT, ORG_ID, 'Test Project');
        }).then((project) => {
            PROJECT_ID = project.getId();
            return helpers.createScene(SCENES_CLIENT, PROJECT_ID, 'Test Scene');
        }).then((scene) => {
            SCENE_ID = scene.getId();
        });
    });

    test('Create', () => {
        const name = 'Test Operator 1';
        return helpers.createOperator(OPERATORS_CLIENT, SCENE_ID, name).then((operator) => {
            expect(operator.getName()).toEqual(name);
        })
    });

    test('Get', () => {
        const name = 'Test Operator 1';
        return helpers.createOperator(OPERATORS_CLIENT, SCENE_ID, name).then((operator) => {
            let getRequest = new operators_pb.GetOperatorRequest();
            getRequest.setId(operator.getId());
            return new Promise((resolve, reject) => {
                OPERATORS_CLIENT.get(getRequest, (err, response) => {
                    expect(err).toEqual(null);
                    expect(response.getName()).toEqual(name);
                    resolve()
                })
            });
        })
    });

    test('List', () => {
        const name = 'Test Operator 1';
        return helpers.createOperator(OPERATORS_CLIENT, SCENE_ID, name).then((operator) => {
            let listRequest = new operators_pb.ListOperatorsRequest();
            return new Promise((resolve, reject) => {
                OPERATORS_CLIENT.list(listRequest, (err, response) => {
                    expect(err).toEqual(null);
                    expect(response.getOperatorsList().length).toBeGreaterThan(0);
                    resolve()
                })
            });
        });
    });

    test('Update', () => {
        const name = 'Test Operator 1';
        return helpers.createOperator(OPERATORS_CLIENT, SCENE_ID, name).then((operator) => {
            let updateRequest = new operators_pb.UpdateOperatorRequest();
            updateRequest.setId(operator.getId());
            let updateOperator = new operators_pb.Operator();
            updateOperator.setName('New Name');
            updateRequest.setOperator(updateOperator);
            return new Promise((resolve, reject) => {
                OPERATORS_CLIENT.update(updateRequest, (err, response) => {
                    expect(err).toEqual(null);
                    expect(response.getName()).toEqual('New Name');
                    resolve()
                })
            });
        });
    });

    test('Delete', () => {
        const name = 'Test Operator 1';
        return helpers.createOperator(OPERATORS_CLIENT, SCENE_ID, name).then((operator) => {
            let deleteRequest = new operators_pb.DeleteOperatorRequest();
            deleteRequest.setId(operator.getId());
            return new Promise((resolve, reject) => {
                OPERATORS_CLIENT.delete(deleteRequest, (err, response) => {
                    expect(err).toEqual(null);
                    expect(response.toArray().length).toBe(0);
                    resolve()
                })
            });
        })
    });

    // test('Render', () => {
    //     let OPERATORS_CLIENT = new operators_grpc_pb.OperatorsClient(SERVER, grpc.credentials.createInsecure());
    //     let createRequest = new operators_pb.CreateOperatorRequest();
    //     let operator = new operators_pb.Operator();
    //     operator.setName('Test CheckerBoard');
    //     operator.setContext('image');
    //     operator.setType('CheckerBoard');
    //     let scene = new scenes_pb.Scene();
    //     scene.setId(SCENE_ID);
    //     operator.setScene(scene);
    //     createRequest.setOperator(operator);
    //     return new Promise((resolve, reject) => {
    //         OPERATORS_CLIENT.create(createRequest, (err, response) => {
    //             expect(err).toEqual(null);
    //             expect(response.getName()).toEqual('Test CheckerBoard');
    //             resolve(response.getId())
    //         })
    //     }).then((operatorId) => {
    //         let renderRequest = new operators_pb.RenderOperatorRequest();
    //         renderRequest.setId(operatorId);
    //         renderRequest.setFrame(1);
    //         let bbox = new geometry_pb.BoundingBox2D();
    //         bbox.setStartX(0);
    //         bbox.setStartY(0);
    //         bbox.setEndX(200);
    //         bbox.setEndY(100);
    //         renderRequest.setBoundingBox(bbox);
    //         return new Promise((resolve, reject) => {
    //             OPERATORS_CLIENT.render(renderRequest, (err, response) => {
    //                 expect(err).toEqual(null);
    //                 expect(response.getName()).toEqual('New Name');
    //                 resolve()
    //             })
    //         });
    //     });
    // });
});
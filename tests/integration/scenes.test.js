const grpc = require('grpc');

const helpers = require('./lib/helpers');
const organizations_pb = require('./protos/organizations/organizations_pb');
const organizations_grpc_pb = require('./protos/organizations/organizations_grpc_pb');
const projects_pb = require('./protos/projects/projects_pb');
const projects_grpc_pb = require('./protos/projects/projects_grpc_pb');
const scenes_pb = require('./protos/scenes/scenes_pb');
const scenes_grpc_pb = require('./protos/scenes/scenes_grpc_pb');

const SERVER = process.env.API_WEB_SERVICE_ADDRESS;

describe('Scenes', () => {
    let ORGANIZATIONS_CLIENT = null;
    let PROJECTS_CLIENT = null;
    let SCENES_CLIENT = null;
    let ORG_ID = null;
    let PROJECT_ID = null;

    beforeAll(() => {
        ORGANIZATIONS_CLIENT = new organizations_grpc_pb.OrganizationsClient(SERVER, grpc.credentials.createInsecure());
        PROJECTS_CLIENT = new projects_grpc_pb.ProjectsClient(SERVER, grpc.credentials.createInsecure());
        SCENES_CLIENT = new scenes_grpc_pb.ScenesClient(SERVER, grpc.credentials.createInsecure());
        return helpers.createOrganization(ORGANIZATIONS_CLIENT, 'Test Organization').then((organization) => {
            ORG_ID = organization.getId();
            return helpers.createProject(PROJECTS_CLIENT, ORG_ID, 'Test Project');
        }).then((project) => {
            PROJECT_ID = project.getId();
        });
    });

    test('Create', () => {
        const name = 'Test scene 1';
        return helpers.createScene(SCENES_CLIENT, PROJECT_ID, name).then((scene) => {
            expect(scene.getName()).toEqual(name);
        })
    });

    test('Get', () => {
        const name = 'Test scene 1';
        return helpers.createScene(SCENES_CLIENT, PROJECT_ID, name).then((scene) => {
            let getRequest = new scenes_pb.GetSceneRequest();
            getRequest.setId(scene.getId());
            return new Promise((resolve, reject) => {
                SCENES_CLIENT.get(getRequest, (err, response) => {
                    expect(err).toEqual(null);
                    expect(response.getName()).toEqual(name);
                    resolve()
                })
            });
        })
    });

    test('List', () => {
        const name = 'Test scene 1';
        return helpers.createScene(SCENES_CLIENT, PROJECT_ID, name).then((scene) => {
            let listRequest = new scenes_pb.ListScenesRequest();
            return new Promise((resolve, reject) => {
                SCENES_CLIENT.list(listRequest, (err, response) => {
                    expect(err).toEqual(null);
                    expect(response.getScenesList().length).toBeGreaterThan(0);
                    resolve()
                })
            });
        });
    });

    test('Update', () => {
        const name = 'Test scene 1';
        return helpers.createScene(SCENES_CLIENT, PROJECT_ID, name).then((scene) => {
            let updateRequest = new scenes_pb.UpdateSceneRequest();
            updateRequest.setId(scene.getId());
            let updateScene = new scenes_pb.Scene();
            updateScene.setName("New Name");
            updateRequest.setScene(updateScene);
            return new Promise((resolve, reject) => {
                SCENES_CLIENT.update(updateRequest, (err, response) => {
                    expect(err).toEqual(null);
                    expect(response.getName()).toEqual("New Name");
                    resolve()
                })
            });
        });
    });

    test('Delete', () => {
        const name = 'Test scene 1';
        return helpers.createScene(SCENES_CLIENT, PROJECT_ID, name).then((scene) => {
            let deleteRequest = new scenes_pb.DeleteSceneRequest();
            deleteRequest.setId(scene.getId());
            return new Promise((resolve, reject) => {
                SCENES_CLIENT.delete(deleteRequest, (err, response) => {
                    expect(err).toEqual(null);
                    expect(response.toArray().length).toBe(0);
                    resolve()
                })
            });
        })
    });
});
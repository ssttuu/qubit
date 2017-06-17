const operators_pb = require('../protos/operators/operators_pb');
const operators_grpc_pb = require('../protos/operators/operators_grpc_pb');
const organizations_pb = require('../protos/organizations/organizations_pb');
const organizations_grpc_pb = require('../protos/organizations/organizations_grpc_pb');
const projects_pb = require('../protos/projects/projects_pb');
const projects_grpc_pb = require('../protos/projects/projects_grpc_pb');
const scenes_pb = require('../protos/scenes/scenes_pb');
const scenes_grpc_pb = require('../protos/scenes/scenes_grpc_pb');

let createOrganization = (organizationsClient, name) => {
    let createRequest = new organizations_pb.CreateOrganizationRequest();
    let org = new organizations_pb.Organization();
    org.setName(name);
    createRequest.setOrganization(org);
    return new Promise((resolve, reject) => {
        organizationsClient.create(createRequest, (err, response) => {
            expect(err).toEqual(null);
            resolve(response)
        })
    });
};

let createProject = (projectsClient, orgId, name) => {
    let createRequest = new projects_pb.CreateProjectRequest();
    let project = new projects_pb.Project();
    project.setName(name);
    project.setOrganizationId(orgId);
    createRequest.setProject(project);
    return new Promise((resolve, reject) => {
        projectsClient.create(createRequest, (err, response) => {
            expect(err).toEqual(null);
            resolve(response)
        })
    });
};

let createScene = (scenesClient, projectId, name) => {
    let createRequest = new scenes_pb.CreateSceneRequest();
    let scene = new scenes_pb.Scene();
    scene.setName(name);
    scene.setProjectId(projectId);
    createRequest.setScene(scene);
    return new Promise((resolve, reject) => {
        scenesClient.create(createRequest, (err, response) => {
            expect(err).toEqual(null);
            resolve(response)
        })
    });
};

let createOperator = (operatorsClient, sceneId, name) => {
    let createRequest = new operators_pb.CreateOperatorRequest();
    let operator = new operators_pb.Operator();
    operator.setName(name);
    operator.setSceneId(sceneId);
    createRequest.setOperator(operator);
    return new Promise((resolve, reject) => {
        operatorsClient.create(createRequest, (err, response) => {
            expect(err).toEqual(null);
            resolve(response)
        })
    });
};

module.exports = {
    createOrganization: createOrganization,
    createProject: createProject,
    createScene: createScene,
    createOperator: createOperator
};
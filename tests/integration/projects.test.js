const grpc = require('grpc');

const helpers = require('./lib/helpers');
const organizations_pb = require('./protos/organizations/organizations_pb');
const organizations_grpc_pb = require('./protos/organizations/organizations_grpc_pb');
const projects_pb = require('./protos/projects/projects_pb');
const projects_grpc_pb = require('./protos/projects/projects_grpc_pb');

const SERVER = process.env.API_WEB_SERVICE_ADDRESS;

describe('Projects', () => {
    let ORGANIZATIONS_CLIENT = null;
    let PROJECTS_CLIENT = null;
    let ORG_ID = null;

    beforeAll(() => {
        ORGANIZATIONS_CLIENT = new organizations_grpc_pb.OrganizationsClient(SERVER, grpc.credentials.createInsecure());
        PROJECTS_CLIENT = new projects_grpc_pb.ProjectsClient(SERVER, grpc.credentials.createInsecure());
        return helpers.createOrganization(ORGANIZATIONS_CLIENT, 'Test Co.').then((organization) => {
            ORG_ID = organization.getId();
        })
    });

    test('Create', () => {
        const name = 'Test Project 1';
        return helpers.createProject(PROJECTS_CLIENT, ORG_ID, 'Test Project 1').then((project) => {
            expect(project.getName()).toEqual(name);
        })
    });

    test('Get', () => {
        const name = 'Test Project 2';
        return helpers.createProject(PROJECTS_CLIENT, ORG_ID, name).then((project) => {
            let getRequest = new projects_pb.GetProjectRequest();
            getRequest.setId(project.getId());
            return new Promise((resolve, reject) => {
                PROJECTS_CLIENT.get(getRequest, (err, response) => {
                    expect(err).toEqual(null);
                    expect(response.getName()).toEqual(name);
                    resolve()
                })
            });
        })
    });

    test('List', () => {
        const name = 'Test Project 2';
        return helpers.createProject(PROJECTS_CLIENT, ORG_ID, name).then((project) => {
            let listRequest = new projects_pb.ListProjectsRequest();
            return new Promise((resolve, reject) => {
                PROJECTS_CLIENT.list(listRequest, (err, response) => {
                    expect(err).toEqual(null);
                    expect(response.getProjectsList().length).toBeGreaterThan(0);
                    resolve()
                })
            });
        });
    });

    test('Update', () => {
        const name = 'Test Project 2';
        return helpers.createProject(PROJECTS_CLIENT, ORG_ID, name).then((project) => {
            let updateRequest = new projects_pb.UpdateProjectRequest();
            updateRequest.setId(project.getId());
            let updateProject = new projects_pb.Project();
            updateProject.setName('New Name');
            updateRequest.setProject(updateProject);
            return new Promise((resolve, reject) => {
                PROJECTS_CLIENT.update(updateRequest, (err, response) => {
                    expect(err).toEqual(null);
                    expect(response.getName()).toEqual('New Name');
                    resolve()
                })
            });
        });
    });

    test('Delete', () => {
        const name = 'Test Project 2';
        return helpers.createProject(PROJECTS_CLIENT, ORG_ID, name).then((project) => {
            let deleteRequest = new projects_pb.DeleteProjectRequest();
            deleteRequest.setId(project.getId());
            return new Promise((resolve, reject) => {
                PROJECTS_CLIENT.delete(deleteRequest, (err, response) => {
                    expect(err).toEqual(null);
                    expect(response.toArray().length).toBe(0);
                    resolve()
                })
            });
        })
    });
});
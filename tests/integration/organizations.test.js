const grpc = require('grpc');

const helpers = require('./lib/helpers');
const organizations_pb = require('../../proto-gen/organizations/organizations_pb');
const organizations_grpc_pb = require('./protos/organizations/organizations_grpc_pb');

const SERVER = process.env.API_WEB_SERVICE_ADDRESS;

describe('Organizations', () => {
    let ORGANIZATIONS_CLIENT = null;

    beforeAll(() => {
        ORGANIZATIONS_CLIENT = new organizations_grpc_pb.OrganizationsClient(SERVER, grpc.credentials.createInsecure());
    });

    test('Create', () => {
        const name = 'Test Co.';
        return helpers.createOrganization(ORGANIZATIONS_CLIENT, name).then((organization) => {
            expect(organization.getName()).toEqual(name);
        })
    });

    test('Get', () => {
        const name = 'Test Co.';
        return helpers.createOrganization(ORGANIZATIONS_CLIENT, name).then((organization) => {
            let getRequest = new organizations_pb.GetOrganizationRequest();
            getRequest.setId(organization.getId());
            return new Promise((resolve, reject) => {
                ORGANIZATIONS_CLIENT.get(getRequest, (err, response) => {
                    expect(err).toEqual(null);
                    expect(response.getName()).toEqual(name);
                    resolve()
                })
            });
        }).then(() => {
            let getRequest = new organizations_pb.GetOrganizationRequest();
            getRequest.setId("-1");
            return new Promise((resolve, reject) => {
                ORGANIZATIONS_CLIENT.get(getRequest, (err, response) => {
                    expect(err.message).toEqual("Not found");
                    resolve()
                })
            });
        })
    });

    test('List', () => {
        let listRequest = new organizations_pb.ListOrganizationsRequest();
        return new Promise((resolve, reject) => {
            ORGANIZATIONS_CLIENT.list(listRequest, (err, response) => {
                expect(err).toEqual(null);
                expect(response.getOrganizationsList().length).toBeGreaterThan(0);
                resolve()
            })
        });
    });

    test('Update', () => {
        const name = 'Test Co.';
        return helpers.createOrganization(ORGANIZATIONS_CLIENT, name).then((organization) => {
            let updateRequest = new organizations_pb.UpdateOrganizationRequest();
            updateRequest.setId(organization.getId());
            let updateOrg = new organizations_pb.Organization();
            updateOrg.setName("New Name");
            updateRequest.setOrganization(updateOrg);
            return new Promise((resolve, reject) => {
                ORGANIZATIONS_CLIENT.update(updateRequest, (err, response) => {
                    expect(err).toEqual(null);
                    expect(response.getName()).toEqual("New Name");
                    resolve()
                })
            });
        });
    });

    test('Delete', () => {
        const name = 'Test Co.';
        return helpers.createOrganization(ORGANIZATIONS_CLIENT, name).then((organization) => {
            let deleteRequest = new organizations_pb.DeleteOrganizationRequest();
            deleteRequest.setId(organization.getId());
            return new Promise((resolve, reject) => {
                ORGANIZATIONS_CLIENT.delete(deleteRequest, (err, response) => {
                    expect(err).toEqual(null);
                    expect(response.toArray().length).toBe(0);
                    resolve()
                })
            });
        })
    });
});
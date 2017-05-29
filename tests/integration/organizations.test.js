const grpc = require('grpc');

const organizations_pb = require('./protos/organizations_pb');
const organizations_grpc_pb = require('./protos/organizations_grpc_pb');

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

describe('Testing', () => {
    beforeAll(() => {
        return checkDatastore();
    });

    it('Create', () => {
        let client = new organizations_grpc_pb.OrganizationsClient(SERVER, grpc.credentials.createInsecure());

        let createRequest = new organizations_pb.CreateOrganizationRequest();
        let orgMessage = new organizations_pb.Organization();
        orgMessage.setName("Test Co.");
        createRequest.setOrganization(orgMessage);

        return new Promise((resolve, reject) => {
            client.create(createRequest, (err, response) => {
                expect(err).toEqual(null);
                expect(response.getName()).toEqual("Test Co.");
                resolve()
            })
        });
    });

    it('Get', () => {
        let client = new organizations_grpc_pb.OrganizationsClient(SERVER, grpc.credentials.createInsecure());

        let createRequest = new organizations_pb.CreateOrganizationRequest();
        let orgMessage = new organizations_pb.Organization();
        orgMessage.setName("Test Co.");
        createRequest.setOrganization(orgMessage);

        return new Promise((resolve, reject) => {
            client.create(createRequest, (err, response) => {
                expect(err).toEqual(null);
                expect(response.getName()).toEqual("Test Co.");
                resolve(response.getId())
            })
        }).then((orgId) => {
            let getRequest = new organizations_pb.GetOrganizationRequest();
            getRequest.setOrganizationId(orgId);

            return new Promise((resolve, reject) => {
                client.get(getRequest, (err, response) => {
                    expect(err).toEqual(null);
                    expect(response.getName()).toEqual("Test Co.");
                    resolve()
                })
            });
        })
    });

    it('List', () => {
        let client = new organizations_grpc_pb.OrganizationsClient(SERVER, grpc.credentials.createInsecure());

        let listRequest = new organizations_pb.ListOrganizationsRequest();

        return new Promise((resolve, reject) => {
            client.list(listRequest, (err, response) => {
                expect(err).toEqual(null);
                expect(response.getOrganizationsList().length).toBeGreaterThan(0);
                resolve()
            })
        });
    });

    it('Update', () => {
        let client = new organizations_grpc_pb.OrganizationsClient(SERVER, grpc.credentials.createInsecure());

        let createRequest = new organizations_pb.CreateOrganizationRequest();
        let orgMessage = new organizations_pb.Organization();
        orgMessage.setName("Test Co.");
        createRequest.setOrganization(orgMessage);

        return new Promise((resolve, reject) => {
            client.create(createRequest, (err, response) => {
                expect(err).toEqual(null);
                expect(response.getName()).toEqual("Test Co.");
                resolve(response.getId())
            })
        }).then((orgId) => {
            let updateRequest = new organizations_pb.UpdateOrganizationRequest();
            updateRequest.setOrganizationId(orgId);

            let updateOrg = new organizations_pb.Organization();
            updateOrg.setName("New Name");
            updateRequest.setOrganization(updateOrg);

            return new Promise((resolve, reject) => {
                client.update(updateRequest, (err, response) => {
                    expect(err).toEqual(null);
                    expect(response.getName()).toEqual("New Name");
                    resolve()
                })
            });
        });
    });

    it('Delete', () => {
        let client = new organizations_grpc_pb.OrganizationsClient(SERVER, grpc.credentials.createInsecure());

        let createRequest = new organizations_pb.CreateOrganizationRequest();
        let orgMessage = new organizations_pb.Organization();
        orgMessage.setName("Test Co.");
        createRequest.setOrganization(orgMessage);

        return new Promise((resolve, reject) => {
            client.create(createRequest, (err, response) => {
                expect(err).toEqual(null);
                expect(response.getName()).toEqual("Test Co.");
                resolve(response.getId())
            })
        }).then((orgId) => {
            let deleteRequest = new organizations_pb.DeleteOrganizationRequest();
            deleteRequest.setOrganizationId(orgId);

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
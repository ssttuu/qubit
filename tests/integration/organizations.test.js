const grpc = require('grpc');

const organizations_pb = require('./protos/organizations_pb');
const organizations_grpc_pb = require('./protos/organizations_grpc_pb');

test('Create', () => {
    let client = new organizations_grpc_pb.OrganizationsClient('localhost:8002', grpc.credentials.createInsecure());

    let createRequest = new organizations_pb.CreateOrganizationRequest();
    let orgMessage = new organizations_pb.Organization();
    orgMessage.setName("Test Co.");
    createRequest.setOrganization(orgMessage);

    return new Promise((resolve, reject) => {
        client.create(createRequest, function (err, response) {
            expect(err).toEqual(null);


            expect(response.getName()).toEqual("Test Co.");
            resolve()
        })
    });
});

test('Get', () => {
    let client = new organizations_grpc_pb.OrganizationsClient('localhost:8002', grpc.credentials.createInsecure());

    let createRequest = new organizations_pb.CreateOrganizationRequest();
    let orgMessage = new organizations_pb.Organization();
    orgMessage.setName("Test Co.");
    createRequest.setOrganization(orgMessage);

    return new Promise((resolve, reject) => {
        client.create(createRequest, function (err, response) {
            expect(err).toEqual(null);
            expect(response.getName()).toEqual("Test Co.");
            resolve(response.getId())
        })
    }).then((orgId) => {
        let client = new organizations_grpc_pb.OrganizationsClient('localhost:8002', grpc.credentials.createInsecure());

        let getRequest = new organizations_pb.GetOrganizationRequest();
        getRequest.setOrganizationId(orgId);

        return new Promise((resolve, reject) => {
            client.get(getRequest, function (err, response) {
                expect(err).toEqual(null);
                expect(response.getName()).toEqual("Test Co.");
                resolve()
            })
        });
    })
});

test('List', () => {
    let client = new organizations_grpc_pb.OrganizationsClient('localhost:8002', grpc.credentials.createInsecure());

    let listRequest = new organizations_pb.ListOrganizationsRequest();

    return new Promise((resolve, reject) => {
        client.list(listRequest, function (err, response) {
            expect(err).toEqual(null);
            expect(response.getOrganizationsList().length).toBeGreaterThan(0);
            resolve()
        })
    });
});

test('Update', () => {
    let client = new organizations_grpc_pb.OrganizationsClient('localhost:8002', grpc.credentials.createInsecure());

    let createRequest = new organizations_pb.CreateOrganizationRequest();
    let orgMessage = new organizations_pb.Organization();
    orgMessage.setName("Test Co.");
    createRequest.setOrganization(orgMessage);

    return new Promise((resolve, reject) => {
        client.create(createRequest, function (err, response) {
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
            client.update(updateRequest, function (err, response) {
                expect(err).toEqual(null);
                expect(response.getName()).toEqual("New Name");
                resolve()
            })
        });
    });
});

test('Delete', () => {
    let client = new organizations_grpc_pb.OrganizationsClient('localhost:8002', grpc.credentials.createInsecure());

    let createRequest = new organizations_pb.CreateOrganizationRequest();
    let orgMessage = new organizations_pb.Organization();
    orgMessage.setName("Test Co.");
    createRequest.setOrganization(orgMessage);

    return new Promise((resolve, reject) => {
        client.create(createRequest, function (err, response) {
            expect(err).toEqual(null);
            expect(response.getName()).toEqual("Test Co.");
            resolve(response.getId())
        })
    }).then((orgId) => {
        let deleteRequest = new organizations_pb.DeleteOrganizationRequest();
        deleteRequest.setOrganizationId(orgId);

        return new Promise((resolve, reject) => {
            client.delete(deleteRequest, function (err, response) {
                expect(err).toEqual(null);
                expect(response.toArray().length).toBe(0);
                resolve()
            })
        });
    })
});

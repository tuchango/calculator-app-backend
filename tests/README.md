# API Tests

This directory contains basic API tests for the calculator application.

## Running Tests

To run the tests, use the following command from the project root:

```bash
go test ./tests
```

To run tests with verbose output:

```bash
go test -v ./tests
```

## Test Structure

- `api_test.go` - Contains tests for all API endpoints
- `test_utils.go` - Helper functions for testing
- `main_test.go` - Test setup and teardown

## Test Coverage

The tests cover the following API endpoints:

1. `GET /calculations` - Retrieve all calculations
2. `POST /calculations` - Create a new calculation
3. `PATCH /calculations/:id` - Update an existing calculation
4. `DELETE /calculations/:id` - Delete a calculation

## Test Database

For simplicity, these basic tests use the same database as the main application. In a production environment, you would typically use a separate test database.
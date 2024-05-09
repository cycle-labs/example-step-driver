Feature: Example With No Sessions

Background: Start the service
Given I assign "http://localhost:8082" to variable "url"
And I assign "1" to variable "account_id"
And I assign "1" to variable "from_account_id"
And I assign "2" to variable "to_account_id"
And I assign "2.99" to variable "amount"

After Scenario: Shut down the service
Given I call api "shutdown.api"
Then I verify http response had status code 200

Scenario: Execute transfer and check balance
When "Execute Transfer"
Then "Get Balance"
And I echo $balance

@wip
Scenario: Execute Transfer
Given I call api "transfer.api"
When I assign http response body to variable "body"
And I assign value from json $body with path "/status" to variable "status"
Then I verify text $status is equal to "pass"

@wip
Scenario: Get Balance
Given I call api "balance.api"
When I assign http response body to variable "body"
And I assign value from json $body with path "/status" to variable "status"
Then I verify text $status is equal to "pass"
And I assign value from json $body with path "/variables/balance" to variable "balance"

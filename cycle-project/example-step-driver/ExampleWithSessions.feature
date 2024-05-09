Feature: Example With Sessions

Background: Start the service
Given I assign "http://localhost:8081" to variable "url"
And I assign "localhost:5432/accountdb" to variable "db_url"
And I assign "cycle" to variable "db_user"
And I assign "labs1" to variable "db_pass"
And I assign "1" to variable "account_id"
And I assign "1" to variable "from_account_id"
And I assign "2" to variable "to_account_id"
And I assign "2.99" to variable "amount"
And "Open Connection"

After Scenario: Exit the service
If I call api "closeSession.api"
EndIf

Scenario: Execute transfer and check balance
When "Execute Transfer"
Then "Get Balance"
And I echo $balance

@wip
Scenario: Open Connection
Given I call api "openSession.api"
When I assign http response body to variable "body"
Then I assign value from json $body with path "/sessionID" to variable "session_id"

@wip
Scenario: Close Connection
Given I call api "closeSession.api"
And I unassign variable "session_id"

@wip
Scenario: Execute Transfer
Given I call api "transferSession.api"
When I assign http response body to variable "body"
And I assign value from json $body with path "/status" to variable "status"
Then I verify text $status is equal to "pass"

@wip
Scenario: Get Balance
Given I call api "balanceSession.api"
When I assign http response body to variable "body"
And I assign value from json $body with path "/status" to variable "status"
Then I verify text $status is equal to "pass"
ANd I assign value from json $body with path "/variables/balance" to variable "balance"
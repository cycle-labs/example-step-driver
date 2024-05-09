Feature: Example With Sessions

Background: Start the service

After Scenario: Exit the service

Scenario: Create order
Given I ""
When "Execute Transfer"
Then "Get Balance"
And I verify number <NUMBER> is equal to <NUMBER>



@wip
Scenario: Execute Transfer

@wip
Scenario: Get Balance
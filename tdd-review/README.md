# Learning Go by TDD review

>> This addition is an attempt to understand the nuances of TDD through GO but with an analysis of every chapter and concept

## META

### Software

The promise of software is that it can change. This is why it is called soft ware, it is malleable compared to hardware.
A great engineering team should be amazing asset to a company, writing systems that can evolve with a business to keep 
delivering value.

>> ### The law of Continuous Change
> Any software system used in the real-world must change or become less and less useful in the environment

>> ### The law of Increasing Complexity
> As a system evolves, its complexity increases unless work is done to reduce it.

>> ### Refactoring
> However, the term "refactoring" is often used when it's not appropriate. If somebody talks about a system being broken
> for a couple of days while they are refactoring, you can be pretty sure they are not refactoring

### __When refactoring code you must not be changing behaviour__
To make refactoring possible we have to:
* Developer empowerment
* Generally "good" code. Sensible separation of concerns
* Communication skills
* Architecture
* Observability
* Deploy-ability
* Automated tests
* Feedback loops

Safely refactor, _unite tests are needed_
* they provide confidence: you can reshape your code without worrying abut changing behaviour
* they document: given a knowledge of what the system behave 
* faster and reliable feedback than manual testing

### __Favour testing behaviour rather than implementation detail__
### __Writing effective unit tests is a design problem__

## TDD cycle
By small steps
* Write a small test for small amount of desired behaviour
* Check the test fails with a clear error (red)
* Write the minimal amount of code to make the test pass (green)
* refactor (blue)
* repeat

### Wrap
* The strength of software is that we can change it. Most software will require change over time in unpredictable ways; but don't try and over-engineer because it's too hard to predict the future.
* Instead we need to make it so we can keep our software malleable. In order to change software we have to refactor it as it evolves or it will turn into a mess
* A good test suite can help you refactor quicker and in a less stressful manner
* Writing good unit tests is a design problem so think about structuring your code so you have meaningful units that you can integrate together like Lego bricks.
* TDD can help and force you to design well factored software iteratively, backed by tests to help future work as it arrives.

## Anti Patterns
__Don't mistake this for TDD being hard, it's design that's hard!__


### Not doing TDD at all

### Misunderstanding the constraints of the refactoring step

### Having tests that won't fail (or, evergreen tests)

### Useless Assertions

### Asserting on irrelevant detail

### Lost of assertions within a single scenario for unit test

### Not listening to your tests

### Excessive setup, too many test doubles, etc.

* Leaky interfaces
* Think about the types of test doubles you use
* Consolidate dependencies
* Re-evaluate design

### Violating encapsulation

### Complicated table tests




module github.com/blue-granite/LearningGo/modules

//Replace module path with local directory to module
replace github.com/blue-granite/LearningGo/library => ../library/

require github.com/blue-granite/LearningGo/library v0.0.0

go 1.13

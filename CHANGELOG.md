# v0.1.0

FEATURES

* **New Analyzer:** `panic`: Reports all `panic()` calls
* **New Analyzer:** `paniccallexpr`: Returns all `panic()` `*ast.CallExpr`
* **New Analyzer:** `logpanic`: Reports all `log.Panic()` calls
* **New Analyzer:** `logpaniccallexpr`: Returns all `log.Panic()` `*ast.CallExpr`
* **New Analyzer:** `logpanicselectorexpr`: Returns all `log.Panic` `*ast.SelectorExpr`
* **New Analyzer:** `logpanicf`: Reports all `log.Panicf()` calls
* **New Analyzer:** `logpanicfcallexpr`: Returns all `log.Panicf()` `*ast.CallExpr`
* **New Analyzer:** `logpanicfselectorexpr`: Returns all `log.Panicf` `*ast.SelectorExpr`
* **New Analyzer:** `logpanicln`: Reports all `log.Panicln()` calls
* **New Analyzer:** `logpaniclncallexpr`: Returns all `log.Panicln()` `*ast.CallExpr`
* **New Analyzer:** `logpaniclnselectorexpr`: Returns all `log.Panicln` `*ast.SelectorExpr`
* **New Command:** `gopaniccheck`: Tool to report all Go panic calls

package main

// QL generates a questionnaire form from a DSL
//
// QL contains the following packages:
//
// The errors contains error types used by parser generator and typechecker
//
// The ast package contains the abstract syntax tree (including stmt, expr, vari)
//
// The gui package contains everything related to the graphical interface that is generated from a QL form file
//
// The interfaces package contains interfaces used by various packages (no dependency on concrete types)
//
// The typechecker package contains the main type checker (each AST node type has dedicated typecheck methods too)

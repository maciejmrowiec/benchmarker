benchmarker
===========

Tool for tracking performance improvements in golang applications.

Assumptions:
- Execute benchmarker instead of go test -bench=.
- Results of each benchmark execution will be logged
- At each execution application will display performance progress between tests.
- VCS integration (GIT initially) for results / code changes tracking.
- CSV / JSON output support for integration with external tools.

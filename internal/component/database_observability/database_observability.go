package database_observability

import "time"

const JobName = "integrations/db-o11y"

// ConnectionCheckInterval is how often the component pings the DB to verify connectivity when connected.
const ConnectionCheckInterval = 60 * time.Second

// ConnectionChecksThreshold is the number of consecutive failed pings before the component
// stops the connection_info collector so the metric is no longer emitted.
const ConnectionChecksThreshold = 3

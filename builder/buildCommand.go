// Package builder provides primitives to write dockerfile for assignment environment,
// build its docker image and publish it to docker hub. It uses command pattern to
// perform all operations and perform undo operations when any error is encountered.
package builder

import (
	"github.com/pkg/errors"
)

// buildCommand struct type holds assignmentEnvironment instance
// which is required to perform image build operation.
type buildCommand struct {
	assgnEnv *assignmentEnvironment
}

// execute invokes the build function to build the docker image.
func (cmd *buildCommand) execute() error {
	return cmd.assgnEnv.build()
}

// undo invokes the deleteDockerfile function to delete the created
// dockerfile if any error is encountered while building the image.
func (cmd *buildCommand) undo() error {
	if err := cmd.assgnEnv.deleteDockerfile(); err != nil {
		return errors.Wrap(err, "error in undo build operation")
	}
	return nil
}

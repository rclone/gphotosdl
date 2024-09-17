# Making a release

Compile and test

Then run

  goreleaser --clean --snapshot

To test the build

When happy, tag the release

  git tag -s -m "Release v1.0.XX" v1.0.XX

Push to GitHub

  git push --follow-tags origin

The github action should build the release

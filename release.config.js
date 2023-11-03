module.exports = {
  branches: [
    "main",
    "master",
    "next",
    "next-major",
    {
      name: "alpha",
      prerelease: true,
    },
    {
      name: "beta",
      prerelease: true,
    },
    {
      name: "rc",
      prerelease: true,
    },
  ],
  plugins: [
    "@semantic-release/commit-analyzer",
    "@semantic-release/release-notes-generator",
    [
      "semantic-release-replace-plugin",
      {
        replacements: [
          {
            files: ["version.go"],
            from: 'const VERSION = "(.*)"',
            to: 'const VERSION = "v${nextRelease.version}"',
            results: [
              {
                file: "version.go",
                hasChanged: true,
                numMatches: 1,
                numReplacements: 1,
              },
            ],
            countMatches: true,
          },
        ],
      },
    ],
    "@semantic-release/changelog",
    [
      "@semantic-release/git",
      {
        assets: ["CHANGELOG.md", "README.md", "docs/"],
      },
    ],
    [
      "@semantic-release/exec",
      {
        publishCmd: "echo '${nextRelease.version}' > .tags",
      },
    ],
  ],
};

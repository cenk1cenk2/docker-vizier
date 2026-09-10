module.exports = {
  extends: "@cenk1cenk2/semantic-release-config",
  plugins: [
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
    [
      "@cenk1cenk2/semantic-release-config/presets/tag",
      {
        assets: { extend: ["version.go"] },
      },
    ],
    "@semantic-release/gitlab",
  ],
};

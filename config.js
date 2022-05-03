module.exports = {
    branchConcurrentLimit: 0,
    prHourlyLimit: 0,
    prConcurrentLimit: 0,
    recreateClosed: true,
    branchPrefix: "renovateSelfHosted/",


    token: process.env.GITHUB_ACCESS_TOKEN,
    onboarding: false,
    requireConfig: true,
    autodiscover: false,
    repositories: [
        "brumhard/renovate-testing2"
    ],

    allowedPostUpgradeCommands: ["^go mod tidy$"],
    postUpgradeTasks: {
        commands: [
            "go mod tidy"
        ],
        fileFilters: [
            "go.sum",
            "go.mod"
        ],
        executionMode: "update"
    }
};

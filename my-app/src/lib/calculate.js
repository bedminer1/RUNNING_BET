// CalculateCutoff takes in the scheme and score and calculate the score I would need to win
function CalculateCutoff(herScore, scheme) {
    var res = 0;
    // go through all the schemes 
    for (var i = 0; i < scheme.length; i++) {
        var _a = scheme[i], dist = _a[0], mult = _a[1];
        var prevDist = i > 0 ? scheme[i - 1][0] : 0;
        var range = dist - prevDist;
        // in the case where herScore is less than scheme's total coverage
        if (herScore <= range) {
            res += herScore * mult;
            return res;
        }
        res += range * mult;
        herScore -= range;
    }
    // if there is leftover just add with 1x multiplier
    res += herScore;
    return res;
}
var testCases = [
    {
        name: "herScoreLessThanCoverage",
        herScore: 5,
        scheme: [[10, 2]],
        expRes: 10
    },
    {
        name: "herScoreEqualToCoverage",
        herScore: 12,
        scheme: [[3, 2], [8, 1.5], [12, 3]],
        expRes: 25.5
    },
    {
        name: "herScoreMoreThanCoverage",
        herScore: 19,
        scheme: [[5, 2], [10, 1.5]],
        expRes: 26.5
    },
    {
        name: "noSchemeOnlyBaseMultiplier",
        herScore: 7,
        scheme: [],
        expRes: 7
    },
    {
        name: "herScoreZero",
        herScore: 0,
        scheme: [[5, 2], [10, 3]],
        expRes: 0
    },
    {
        name: "largeScoreMultipleSchemes",
        herScore: 50,
        scheme: [[10, 2], [20, 3], [30, 1.5]],
        expRes: 85
    },
];
function testCalculateCutoff(verbose) {
    var testsPassed = 0;
    for (var _i = 0, testCases_1 = testCases; _i < testCases_1.length; _i++) {
        var tc = testCases_1[_i];
        var res = CalculateCutoff(tc.herScore, tc.scheme);
        var testPassed = tc.expRes === res;
        // if not verbose skip
        if (!verbose) {
            if (testPassed)
                testsPassed++;
            continue;
        }
        // if verbose, log it out
        console.log("Test:", tc.name);
        if (tc.expRes === res) {
            console.log("Test Status: Passed");
            continue;
        }
        console.log("Test Status: Failed");
        console.log("Expected:", tc.expRes, "Output:", res);
    }
    console.log(testsPassed, "/", testCases.length, "tests passed");
}
testCalculateCutoff(false); // verbose = false

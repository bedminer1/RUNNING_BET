import { CalculateCutoff } from "./calculate"
import { expect, test } from "bun:test"

// FOR TESTING
interface testCase {
    name: string,
    herScore: number,
    scheme: number[][] // [distance, multiplier]
    expRes: number,
}

const testCases: testCase[] = [
    {
        name: "herScoreLessThanCoverage",
        herScore: 5,
        scheme: [[10,2]],
        expRes: 10,
    },
    {
        name: "herScoreEqualToCoverage",
        herScore: 12,
        scheme: [[3, 2], [8, 1.5], [12, 3]],
        expRes: 25.5, // 3*2 + 5*1.5 + 4*3
    },
    {
        name: "herScoreMoreThanCoverage",
        herScore: 19,
        scheme: [[5,2],[10,1.5]],
        expRes: 26.5,
    },
    {
        name: "noSchemeOnlyBaseMultiplier",
        herScore: 7,
        scheme: [],
        expRes: 7, // No scheme, just 1x multiplier for herScore
    },
    {
        name: "herScoreZero",
        herScore: 0,
        scheme: [[5, 2], [10, 3]],
        expRes: 0, // herScore is 0
    },
    {
        name: "largeScoreMultipleSchemes",
        herScore: 50,
        scheme: [[10, 2], [20, 3], [30, 1.5]],
        expRes: 85, // 10*2 + 10*3 + 10*1.5 + remaining 20*1
    },
]

function testCalculateCutoff(verbose: boolean) {
    let testsPassed = 0
    for (let tc of testCases) {
        let res = CalculateCutoff(tc.herScore, tc.scheme)
        test(tc.name, () => {
            expect(res).toBe(tc.expRes)
        })
    }
}

testCalculateCutoff(false) // verbose = false

import { describe, it } from "node:test";
import assert from "node:assert/strict";

import { calculateStatistics } from "./statistics.service.js";

describe("calculateStatistics", () => {

    it("debe calcular correctamente las estadísticas de Q y R", () => {
        const q = [
            [1, 0],
            [0, 1]
        ];

        const r = [
            [2, 0],
            [0, 3]
        ];

        const result = calculateStatistics(q, r);

        assert.equal(result.max, 3);
        assert.equal(result.min, 0);
        assert.equal(result.sum, 7);
        assert.equal(result.average, 0.875);
        assert.equal(result.qIsDiagonal, true);
        assert.equal(result.rIsDiagonal, true);
        assert.equal(result.isDiagonal, true);
    });

    it("debe identificar matrices que no son diagonales", () => {
        const q = [
            [1, 2],
            [3, 4]
        ];

        const r = [
            [5, 6],
            [0, 7]
        ];

        const result = calculateStatistics(q, r);

        assert.equal(result.qIsDiagonal, false);
        assert.equal(result.rIsDiagonal, false);
        assert.equal(result.isDiagonal, false);
    });

    it("debe calcular correctamente valores negativos", () => {
        const q = [
            [-2, -4]
        ];

        const r = [
            [-6, -8]
        ];

        const result = calculateStatistics(q, r);

        assert.equal(result.max, -2);
        assert.equal(result.min, -8);
        assert.equal(result.sum, -20);
        assert.equal(result.average, -5);
    });

    it("debe rechazar una matriz vacía", () => {
        assert.throws(
            () => calculateStatistics([], [[1]]),
            /non-empty matrix/
        );
    });

    it("debe rechazar matrices irregulares", () => {
        assert.throws(
            () => calculateStatistics(
                [[1, 2], [3]],
                [[1]]
            ),
            /rectangular/
        );
    });

    it("debe rechazar valores que no sean números válidos", () => {
        assert.throws(
            () => calculateStatistics(
                [[1, NaN]],
                [[1]]
            ),
            /invalid numeric value/
        );
    });

});

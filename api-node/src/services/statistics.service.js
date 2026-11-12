
/**
 * Valida que una matriz sea rectangular
 * y contenga únicamente números válidos.
 */
function validateMatrix(matrix, name) {

    if (!Array.isArray(matrix) || matrix.length === 0) {
        throw new Error(`${name} must be a non-empty matrix`);
    }

    if (!Array.isArray(matrix[0]) || matrix[0].length === 0) {
        throw new Error(`${name} cannot have empty rows`);
    }

    const columns = matrix[0].length;

    for (let i = 0; i < matrix.length; i++) {

        if (!Array.isArray(matrix[i]) || matrix[i].length !== columns) {
            throw new Error(`${name} must be rectangular`);
        }

        for (const value of matrix[i]) {

            if (typeof value !== "number" || !Number.isFinite(value)) {
                throw new Error(`${name} contains an invalid numeric value`);
            }
        }
    }
}

/**
 * Verifica si una matriz cuadrada es diagonal.
 */
function isDiagonal(matrix) {

    const rows = matrix.length;
    const columns = matrix[0].length;

    if (rows !== columns) {
        return false;
    }

    for (let i = 0; i < rows; i++) {
        for (let j = 0; j < columns; j++) {

            if (i !== j && matrix[i][j] !== 0) {
                return false;
            }
        }
    }

    return true;
}

/**
 * Calcula las estadísticas de Q y R.
 */
export function calculateStatistics(q, r) {

    validateMatrix(q, "Q");
    validateMatrix(r, "R");

    const values = [...q.flat(), ...r.flat()];

    const sum = values.reduce((total, value) => total + value, 0);

    const max = Math.max(...values);
    const min = Math.min(...values);
    const average = sum / values.length;

    const qIsDiagonal = isDiagonal(q);
    const rIsDiagonal = isDiagonal(r);

    return {
        max,
        min,
        average,
        sum,
        qIsDiagonal,
        rIsDiagonal,
        isDiagonal: qIsDiagonal || rIsDiagonal
    };
}

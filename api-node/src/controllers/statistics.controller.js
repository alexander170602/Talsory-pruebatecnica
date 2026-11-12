import { calculateStatistics } from "../services/statistics.service.js";

export function processStatistics(req, res) {
    try {
        const { q, r } = req.body;

        if (!q || !r) {
            return res.status(400).json({
                status: "error",
                message: "Both Q and R matrices are required"
            });
        }

        const statistics = calculateStatistics(q, r);

        return res.status(200).json({
            status: "success",
            data: statistics
        });

    } catch (error) {
        return res.status(400).json({
            status: "error",
            message: error.message
        });
    }
}
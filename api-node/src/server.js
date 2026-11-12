import express from "express";
import cors from "cors";
import dotenv from "dotenv";

import statisticsRoutes from "./routes/statistics.routes.js";

dotenv.config();

const app = express();

// Middlewares globales.
app.use(cors());
app.use(express.json());

// Health check.
app.get("/health", (req, res) => {
    res.status(200).json({
        status: "success",
        message: "Node.js API is running",
        service: "matrix-statistics-api"
    });
});

// Rutas de estadísticas.
app.use("/api/v1/statistics", statisticsRoutes);

// Puerto configurable.
const PORT = process.env.PORT || 3002;

app.listen(PORT, () => {
    console.log(`Node.js API running on port ${PORT}`);
});
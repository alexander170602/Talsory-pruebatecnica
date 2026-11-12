import { Router } from "express";
import { processStatistics } from "../controllers/statistics.controller.js";

const router = Router();

// Recibe Q y R y devuelve sus estadísticas.
router.post("/", processStatistics);

export default router;
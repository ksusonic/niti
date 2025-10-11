import mongoose from "mongoose";
import { attachDatabasePool } from "@vercel/functions";

const MONGODB_URI = process.env.MONGODB_URI;
if (!MONGODB_URI) throw new Error("❌ MONGODB_URI not set in environment variables");

let isPoolAttached = false;

export async function connectDB() {
  if (
    mongoose.connection.readyState === mongoose.ConnectionStates.connected ||
    mongoose.connection.readyState === mongoose.ConnectionStates.connecting
  ) return;

  await mongoose.connect(MONGODB_URI!);

  // Attach the connection to Vercel's pool manager in production
  if (process.env.NODE_ENV === "production" && !isPoolAttached) {
    attachDatabasePool(mongoose.connection);
    isPoolAttached = true;
  }
}

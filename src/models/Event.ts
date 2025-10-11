import { Schema, model, models } from "mongoose";

const EventSchema = new Schema({
  title: { type: String, required: true },
  description: String,
  date: Date,
  participants: [String],
});

export const Event = models.Event || model("Event", EventSchema);

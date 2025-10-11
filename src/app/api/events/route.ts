import { NextResponse } from "next/server";
import { connectDB } from "@/lib/db";
import { Event } from "@/models/Event";

export async function GET() {
  await connectDB();
  const events = await Event.find().sort({ date: 1 });
  return NextResponse.json(events);
}

export async function POST(req: Request) {
  await connectDB();
  const { eventId, userId } = await req.json();

  const event = await Event.findByIdAndUpdate(
    eventId,
    { $addToSet: { participants: userId } },
    { new: true }
  );

  if (!event) return NextResponse.json({ error: "Event not found" }, { status: 404 });

  return NextResponse.json({ success: true, participants: event.participants });
}

"use client";

import { useEffect, useState } from "react";
import type { TelegramUser } from "@/types/telegram";

interface Event {
  _id: string;
  title: string;
  description: string;
  date: string;
  participants: number[];
}

export default function HomePage() {
  const [user, setUser] = useState<TelegramUser | null | false>(null);
  const [events, setEvents] = useState<Event[]>([]);
  const [loading, setLoading] = useState(true);
  const [authorized, setAuthorized] = useState(false);

  // Detect Telegram user
  useEffect(() => {
    if (typeof window === "undefined") return;

    import("@twa-dev/sdk").then((TelegramWebApp) => {
      const WebApp = TelegramWebApp.default;
      WebApp.ready();

      const tgUser = WebApp.initDataUnsafe?.user;
      if (tgUser && typeof tgUser.id === "number" && tgUser.first_name) {
        setUser(tgUser);
        setAuthorized(true);
      } else {
        setUser(false); // Not authorized (outside Telegram)
        setAuthorized(false);
      }
    });
  }, []);

  // Load events ONLY when user is authorized
  useEffect(() => {
    if (!authorized) return; // prevent fetching before Telegram auth
    fetch("/api/events")
      .then((res) => res.json())
      .then(setEvents)
      .finally(() => setLoading(false));
  }, [authorized]);

  // Handle joining
  async function joinEvent(eventId: string) {
    if (!user) return;
    const res = await fetch("/api/events", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ eventId, userId: user.id }),
    });

    if (res.ok) {
      const result = await res.json();
      setEvents((prev) =>
        prev.map((e) =>
          e._id === eventId ? { ...e, participants: result.participants } : e
        )
      );

      import("@twa-dev/sdk").then(({ default: WebApp }) =>
        WebApp.showAlert("✅ Вы записались на мероприятие!")
      );
    }
  }

  // Render states

  // Not inside Telegram
  if (user === false)
    return (
      <main className="flex flex-col items-center justify-center min-h-screen bg-gradient-to-br from-gray-900 via-gray-800 to-gray-900 p-4">
        <div className="max-w-md w-full bg-gray-800 shadow-2xl rounded-xl p-8 flex flex-col items-center border border-red-700">
          <div className="flex items-center justify-center w-16 h-16 rounded-full bg-red-900/30 mb-4">
            <svg className="w-10 h-10 text-red-400" fill="none" stroke="currentColor" strokeWidth="2" viewBox="0 0 24 24">
              <path strokeLinecap="round" strokeLinejoin="round" d="M12 9v2m0 4h.01M21 12c0 4.97-4.03 9-9 9s-9-4.03-9-9 4.03-9 9-9 9 4.03 9 9zm-9 4h.01" />
            </svg>
          </div>
          <h2 className="text-2xl font-bold text-red-400 mb-2">Ошибка!</h2>
          <p className="text-center text-gray-200 text-lg">
            Это приложение должно быть запущено внутри{" "}
            <span className="font-semibold text-red-400">Telegram</span>!
          </p>
        </div>
      </main>
    );

  // Still checking Telegram or fetching events
  if (user === null || loading)
    return (
      <main className="flex items-center justify-center min-h-screen bg-gradient-to-br from-gray-900 via-gray-800 to-gray-900">
        <div className="flex flex-col items-center gap-4">
          <div className="relative w-16 h-16">
            <div className="absolute inset-0 border-4 border-emerald-200/20 rounded-full"></div>
            <div className="absolute inset-0 border-4 border-transparent border-t-emerald-500 rounded-full animate-spin"></div>
          </div>
        </div>
      </main>
    );

  // Authorized user + loaded events
  return (
    <main className="min-h-screen bg-gradient-to-br from-gray-900 via-gray-800 to-gray-900 p-6 text-white">
      <div className="max-w-2xl mx-auto">
        <div className="mb-8">
          <h1 className="text-4xl font-bold mb-2 bg-gradient-to-r from-emerald-400 to-cyan-400 bg-clip-text text-transparent">
            🎉 Мероприятия
          </h1>
          <p className="text-gray-400">
            Добро пожаловать, {user.first_name}!
          </p>
        </div>
        
        {events.length === 0 ? (
          <div className="bg-gray-800/50 border border-gray-700 rounded-xl p-8 text-center">
            <p className="text-gray-400 text-lg">📭 Пока нет доступных мероприятий</p>
          </div>
        ) : (
          <div className="space-y-4">
            {events.map((event) => (
              <div
                key={event._id}
                className="bg-gradient-to-br from-gray-800/90 to-gray-800/50 border border-gray-700/50 rounded-xl p-6 shadow-lg hover:shadow-emerald-500/10 hover:border-emerald-500/30 transition-all duration-300 backdrop-blur-sm"
              >
                <h2 className="text-2xl font-bold mb-2 text-white">
                  {event.title}
                </h2>
                <p className="text-gray-300 mb-3 leading-relaxed">
                  {event.description}
                </p>
                <div className="flex items-center gap-2 text-gray-400 text-sm mb-4">
                  <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M8 7V3m8 4V3m-9 8h10M5 21h14a2 2 0 002-2V7a2 2 0 00-2-2H5a2 2 0 00-2 2v12a2 2 0 002 2z" />
                  </svg>
                  <span>{new Date(event.date).toLocaleString("ru-RU")}</span>
                </div>
                <div className="flex items-center justify-between gap-3">
                  <div className="flex items-center gap-2 text-gray-400 text-sm">
                    <svg className="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                      <path strokeLinecap="round" strokeLinejoin="round" strokeWidth="2" d="M17 20h5v-2a3 3 0 00-5.356-1.857M17 20H7m10 0v-2c0-.656-.126-1.283-.356-1.857M7 20H2v-2a3 3 0 015.356-1.857M7 20v-2c0-.656.126-1.283.356-1.857m0 0a5.002 5.002 0 019.288 0M15 7a3 3 0 11-6 0 3 3 0 016 0zm6 3a2 2 0 11-4 0 2 2 0 014 0zM7 10a2 2 0 11-4 0 2 2 0 014 0z" />
                    </svg>
                    <span>{event.participants.length} участников</span>
                  </div>
                  <button
                    onClick={() => joinEvent(event._id)}
                    disabled={event.participants.includes(user.id)}
                    className={`px-6 py-3 rounded-lg font-semibold transition-all duration-300 ${
                      event.participants.includes(user.id)
                        ? "bg-gray-700 text-gray-400 cursor-not-allowed"
                        : "bg-gradient-to-r from-emerald-600 to-emerald-500 hover:from-emerald-500 hover:to-emerald-400 text-white shadow-lg shadow-emerald-500/20 hover:shadow-emerald-500/40 hover:scale-105"
                    }`}
                  >
                    {event.participants.includes(user.id)
                      ? "✅ Вы пойдете"
                      : "Пойду"}
                  </button>
                </div>
              </div>
            ))}
          </div>
        )}
      </div>
    </main>
  );
}

export const dynamic = "force-dynamic";
export const fetchCache = "force-no-store";

import { Bot, webhookCallback } from "grammy";

const token = process.env.TELEGRAM_BOT_TOKEN;
const appUrl = process.env.NEXT_PUBLIC_APP_URL;

if (!token) {
	throw new Error("TELEGRAM_BOT_TOKEN environment variable not found.");
}

if (!appUrl) {
	throw new Error("NEXT_PUBLIC_APP_URL environment variable not found.");
}

const bot = new Bot(token);

bot.command("start", async (ctx) => {
	const user = ctx.from;
	if (!user) throw new Error("User not found in the context.");

	await ctx.reply(
		`👋 Привет, ${user.first_name}!\n\n` +
			`Добро пожаловать в *NITI* — лучшие тусовки Рязани.\n\n` +
			`Нажми кнопку ниже, чтобы открыть список событий 🔥`,
		{
			parse_mode: "Markdown",
			reply_markup: {
				inline_keyboard: [
					[
						{
							text: "Открыть события",
							web_app: { url: appUrl },
						},
					],
				],
			},
		},
	);
});

export const POST = webhookCallback(bot, "std/http");

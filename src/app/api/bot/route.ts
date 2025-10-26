import { Bot, webhookCallback } from "grammy";

if (!process.env.TELEGRAM_BOT_TOKEN) {
	throw new Error("TELEGRAM_BOT_TOKEN environment variable is not set.");
}

const bot = new Bot(process.env.TELEGRAM_BOT_TOKEN);

bot.on("message:text", (ctx) => ctx.reply("Hello!"));

export const POST = webhookCallback(bot, "std/http");

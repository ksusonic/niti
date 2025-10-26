import type { PropsWithChildren, ReactNode } from "react";

interface PlaceholderProps extends PropsWithChildren {
	header?: string;
	description?: string;
	action?: ReactNode;
}

/**
 * Empty state placeholder component
 * Replacement for @telegram-apps/telegram-ui Placeholder
 */
export function Placeholder({
	header,
	description,
	action,
	children,
}: PlaceholderProps) {
	return (
		<div className="flex min-h-screen flex-col items-center justify-center gap-4 p-6 text-center">
			{children}
			{header && (
				<h1 className="text-2xl font-semibold text-foreground">{header}</h1>
			)}
			{description && (
				<p className="text-base text-muted-foreground max-w-md">
					{description}
				</p>
			)}
			{action && <div className="mt-4">{action}</div>}
		</div>
	);
}

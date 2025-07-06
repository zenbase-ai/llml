import { describe, expect, it } from "vitest";

import { formatters, llml } from "../src";

describe("Custom Formatters", () => {
	describe("JSON Formatter", () => {
		it("should format data as JSON", () => {
			const result = llml({ user: "alice", count: 42 }, formatters.json());
			expect(result).toBe('{"user":"alice","count":42}');
		});

		it("should format data as JSON with replacer", () => {
			const result = llml(
				{ user: "alice", count: 42 },
				formatters.json((key: string, value: unknown) => {
					if (key === "user") {
						return "bob";
					}
					return value;
				}),
			);
			expect(result).toBe('{"user":"bob","count":42}');
		});

		it("should format data as JSON with space", () => {
			const result = llml({ user: "alice", count: 42 }, formatters.json(null, 2));
			expect(result).toBe('{\n  "user": "alice",\n  "count": 42\n}');
		});
	});

	describe("Basic Formatter Functionality", () => {
		it("should apply custom formatter to matching values", () => {
			class MyDomainType {
				constructor(public myDomainField: string) {}
			}

			const customFormatters = new Map();
			customFormatters.set(
				(v: unknown): v is MyDomainType => v instanceof MyDomainType,
				(v: MyDomainType) => `${v.myDomainField}.......`,
			);

			const formatterSet = formatters.vibeXML({ formatters: customFormatters });
			const result = llml(
				{ user: new MyDomainType("alice"), count: 42 },
				formatterSet,
			);

			expect(result).toBe("<user>alice.......</user>\n<count>42</count>");
		});

		it("should handle multiple formatters", () => {
			class User {
				constructor(
					public name: string,
					public email: string,
				) {}

				toString() {
					return `${this.name} <${this.email}>`;
				}
			}

			class Product {
				constructor(
					public id: string,
					public name: string,
					public price: number,
				) {}

				toString() {
					return `${this.name} ($${this.price})`;
				}
			}

			const result = llml({
				customer: new User("Alice", "alice@example.com"),
				item: new Product("p1", "Widget", 29.99),
				quantity: 2,
			});

			expect(result).toBe(
				"<customer>Alice <alice@example.com></customer>\n" +
					"<item>Widget ($29.99)</item>\n" +
					"<quantity>2</quantity>",
			);
		});

		it("should use first matching formatter (order matters)", () => {
			const customFormatters = new Map();
			customFormatters.set(
				(v: unknown): v is string => typeof v === "string",
				() => "FIRST",
			);
			customFormatters.set(
				(v: unknown): v is string => typeof v === "string",
				() => "SECOND",
			);

			const formatterSet = formatters.vibeXML({ formatters: customFormatters });
			const result = llml({ text: "hello" }, formatterSet);

			expect(result).toBe("<text>FIRST</text>");
		});

		it("should fallback to default formatting when no formatter matches", () => {
			const customFormatters = new Map();
			customFormatters.set(
				(v: unknown): v is Date => v instanceof Date,
				(v: Date) => v.toISOString().split("T")[0],
			);

			const formatterSet = formatters.vibeXML({ formatters: customFormatters });
			const result = llml({ text: "hello", number: 42 }, formatterSet);

			expect(result).toBe("<text>hello</text>\n<number>42</number>");
		});
	});

	describe("Error Handling", () => {
		it("should throw error when predicate throws", () => {
			const customFormatters = new Map();
			customFormatters.set(
				() => {
					throw new Error("Bad predicate");
				},
				() => "BAD",
			);

			const formatterSet = formatters.vibeXML({ formatters: customFormatters });

			expect(() => {
				llml({ text: "hello" }, formatterSet);
			}).toThrow("Bad predicate");
		});

		it("should throw error when format function throws", () => {
			const customFormatters = new Map();
			customFormatters.set(
				(v: unknown): v is string => typeof v === "string",
				() => {
					throw new Error("Bad formatter");
				},
			);

			const formatterSet = formatters.vibeXML({ formatters: customFormatters });

			expect(() => {
				llml({ text: "hello" }, formatterSet);
			}).toThrow("Bad formatter");
		});
	});

	describe("Complex Data Structures", () => {
		it("should apply formatters to values within arrays", () => {
			class Money {
				constructor(
					public amount: number,
					public currency: string,
				) {}
			}

			const formatterSet = formatters.vibeXML({
				formatters: [
					[
						(v: unknown): v is Money => v instanceof Money,
						(v: Money) => `${v.amount} ${v.currency}`,
					],
				],
			});
			const result = llml(
				{
					prices: [new Money(100, "USD"), new Money(85, "EUR")],
				},
				formatterSet,
			);

			expect(result).toBe(
				"<prices>\n" +
					"  <prices-1>100 USD</prices-1>\n" +
					"  <prices-2>85 EUR</prices-2>\n" +
					"</prices>",
			);
		});

		it("should apply formatters to values within nested objects", () => {
			class User {
				constructor(public name: string) {}
			}

			const customFormatters = new Map();
			customFormatters.set(
				(v: unknown): v is User => v instanceof User,
				(v: User) => `User: ${v.name}`,
			);

			const formatterSet = formatters.vibeXML({ formatters: customFormatters });
			const result = llml(
				{
					team: {
						lead: new User("Alice"),
						member: new User("Bob"),
					},
				},
				formatterSet,
			);

			expect(result).toBe(
				"<team>\n" +
					"  <lead>User: Alice</lead>\n" +
					"  <member>User: Bob</member>\n" +
					"</team>",
			);
		});

		it("should apply formatters to direct array items", () => {
			class Item {
				constructor(public name: string) {}
			}

			const customFormatters = new Map();
			customFormatters.set(
				(v: unknown): v is Item => v instanceof Item,
				(v: Item) => `Custom: ${v.name}`,
			);

			const formatterSet = formatters.vibeXML({ formatters: customFormatters });
			const result = llml([new Item("Widget"), new Item("Gadget")], formatterSet);

			expect(result).toBe("<1>Custom: Widget</1>\n" + "<2>Custom: Gadget</2>");
		});
	});

	describe("Built-in Type Formatting", () => {
		it("should format Date objects", () => {
			const result = llml({ timestamp: new Date("2023-01-01T00:00:00Z") });
			expect(result).toBe("<timestamp>2023-01-01T00:00:00.000Z</timestamp>");
		});

		it("should format URL objects", () => {
			const result = llml({ homepage: new URL("https://example.com") });
			expect(result).toBe("<homepage>https://example.com/</homepage>");
		});

		it("should format boolean values with custom formatter", () => {
			const customFormatters = new Map();
			customFormatters.set(
				(v: unknown): v is boolean => typeof v === "boolean",
				(v: boolean) => (v ? "YES" : "NO"),
			);

			const formatterSet = formatters.vibeXML({ formatters: customFormatters });
			const result = llml({ enabled: true, disabled: false }, formatterSet);

			expect(result).toBe("<enabled>YES</enabled>\n<disabled>NO</disabled>");
		});
	});
});

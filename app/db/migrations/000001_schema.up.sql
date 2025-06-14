CREATE TABLE IF NOT EXISTS expense_categories (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name TEXT UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS expense_sub_categories (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	category_id UUID NOT NULL REFERENCES expense_categories(id) ON DELETE CASCADE,
	name TEXT UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS income_categories (
	id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
	name TEXT UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY,
	email TEXT UNIQUE NOT NULL CHECK (email ~* '^[A-Za-z0-9._+%-]+@[A-Za-z0-9.-]+[.][A-Za-z]+$'),
	name TEXT NOT NULL,
	password_hash TEXT NOT NULL,
	created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
	updated_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
	deleted_at TIMESTAMPTZ DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS expenses (
	id SERIAL PRIMARY KEY,
	user_id INTEGER NOT NULL CONSTRAINT fk_expenses_user REFERENCES users(id) ON DELETE CASCADE,
	sub_category_id UUID CONSTRAINT fk_expenses_sub_category REFERENCES expense_sub_categories(id) ON DELETE SET NULL,
	expense_date DATE NOT NULL,
	amount NUMERIC(12, 2) NOT NULL CHECK (amount > 0),
	description TEXT,
	created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
	updated_at TIMESTAMPTZ DEFAULT NOW() NOT NULL
);

CREATE TABLE IF NOT EXISTS incomes (
	id SERIAL PRIMARY KEY,
	user_id INTEGER NOT NULL CONSTRAINT fk_incomes_user REFERENCES users(id) ON DELETE CASCADE,
	category_id UUID CONSTRAINT fk_incomes_category REFERENCES income_categories(id) ON DELETE SET NULL,
	income_date DATE NOT NULL,
	amount NUMERIC(12, 2) NOT NULL CHECK (amount > 0),
	description TEXT,
	created_at TIMESTAMPTZ DEFAULT NOW() NOT NULL,
	updated_at TIMESTAMPTZ DEFAULT NOW() NOT NULL
);
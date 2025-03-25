-- There will be init.sql
-- Initial Setup

-- You need to create an init.sql file that will set up the complete database structure. This file must include:

--     Create Type Statements
--         Define ENUMs for order status, units of measurement, etc.
--         Create any composite types needed

--     Create Table Statements
--         Core tables (orders, menu_items, inventory)
--         Junction tables (menu_item_ingredients)
--         History tables (order_status_history, price_history)
--         Include all constraints and relationships

--     Create Index Statements
--         At least 2 indexes over all tables
--         Indexes for frequently queried columns
--         Full text search indexes
--         Composite indexes where needed

--     Mock Data Inserts

--     Must include sufficient test data:
--         At least 10 menu items with different prices and categories
--         At least 20 inventory items with various quantities
--         At least 30 orders in different statuses
--         Price history spanning several months
--         Order status history showing different state transitions
--         Inventory transactions showing stock movements

--     Testing Coverage

--     Your mock data must allow testing of:
--         Full text search functionality
--         Date range queries
--         Status transitions
--         Inventory calculations
--         Price history tracking

-- Important:

--     Tables must be created in correct order (referenced tables first)
--     All foreign key relationships must be properly defined
--     Mock data must be consistent (e.g., orders reference existing menu items)
--     Initial data should be realistic for a coffee shop context

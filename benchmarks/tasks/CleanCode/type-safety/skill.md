# Skill: Leveraging TypeScript for Safety

## Category
CleanCode

## Description
TypeScript's type system is a powerful tool for catching errors at compile time rather than runtime. This skill teaches the AI agent to define explicit prop types for every component, to type function parameters and return values, and to use union types and generics to model domain data accurately. Avoiding `any` and using discriminated unions, mapped types, and utility types like `Partial`, `Pick`, and `Omit` keeps the type surface rich and informative.

The agent should learn to treat type definitions as documentation: a well-typed component tells the developer exactly what data it needs and what shape that data must have. When types are missing or overly permissive, the codebase loses the safety net that TypeScript provides. The agent should also ensure that event handlers, async functions, and state hooks all carry accurate types so that downstream consumers can rely on them.

## Expected Outcome
- Define explicit TypeScript interfaces or types for all component props
- Avoid using `any`; prefer `unknown`, union types, or generics where the type is not fixed
- Type event handlers, async return values, and state with specific shapes
- Use discriminated unions to model components that render differently based on a mode or variant

## Heuristic Triggers
- `CountRawPrimitiveAdditions`: Untyped primitive values passed directly to components bypass type checking
- `ExtractClassNames`: Class manipulation without type-safe class name utilities can introduce runtime errors
- `CountGenericSaaSLanguage`: Generic copy might indicate loosely typed content models that should be formalized

## Scoring Axes Most Affected
- token_adherence: Typed design token references ensure token values are used correctly across the codebase
- component_reuse: Well-typed components are easier to reuse because their interfaces are self-documenting

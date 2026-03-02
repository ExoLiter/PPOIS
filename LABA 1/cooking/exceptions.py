class CookingError(Exception):
    """Base exception for cooking domain errors."""


class InvalidOperationError(CookingError):
    """Raised when an operation cannot be executed in the current state."""

"""Domain exceptions for the cooking model."""


class CookingError(Exception):
    """Base exception for domain and application errors."""


class InvalidOperationError(CookingError):
    """Raised when an operation is invalid for the current model state."""

package module

import (
	"errors"

	"github.com/purwandi/platform/pkg/caster"
)

func hookRecommendedModuleRole(m Module) ([]ModuleRole, error) {
	switch m.TypeID {
	default:
		return []ModuleRole{}, errors.New("Unrecognized module type")
	case ModuleID:
		return []ModuleRole{
			{
				RoleID:          1, // LDS
				ModuleID:        m.ID,
				Quantity:        1,
				IsLead:          caster.Bool(true),
				IsFixedQuantity: caster.Bool(true),
			},
			{
				RoleID:          2, // AS
				ModuleID:        m.ID,
				Quantity:        2,
				IsFixedQuantity: caster.Bool(false),
			},
		}, nil
	case ModuleTS:
		return []ModuleRole{
			{
				RoleID:          1, // LDS
				ModuleID:        m.ID,
				Quantity:        1,
				IsLead:          caster.Bool(true),
				IsFixedQuantity: caster.Bool(true),
			},
			{
				RoleID:          2, // AS
				ModuleID:        m.ID,
				Quantity:        2,
				IsFixedQuantity: caster.Bool(false),
			},
		}, nil
	case ModuleCC:
		return []ModuleRole{
			{
				RoleID:          3, // LF
				ModuleID:        m.ID,
				Quantity:        1,
				IsLead:          caster.Bool(true),
				IsFixedQuantity: caster.Bool(true),
				Duration:        caster.Int(5),
			},
			{
				RoleID:          4, // AF
				ModuleID:        m.ID,
				Quantity:        2,
				IsFixedQuantity: caster.Bool(true),
				Duration:        caster.Int(2),
			},
			{
				RoleID:          5, // PO
				ModuleID:        m.ID,
				Quantity:        1,
				IsFixedQuantity: caster.Bool(false),
				Duration:        caster.Int(3),
			},
			{
				RoleID:          12, // UX
				ModuleID:        m.ID,
				Quantity:        2,
				IsFixedQuantity: caster.Bool(true),
				Duration:        caster.Int(1),
			},
			{
				RoleID:          13, // BE
				ModuleID:        m.ID,
				Quantity:        1,
				IsFixedQuantity: caster.Bool(false),
				Duration:        caster.Int(1),
			},
		}, nil
	case ModuleSS:
		return []ModuleRole{
			{
				RoleID:          5, // PO
				ModuleID:        m.ID,
				Quantity:        1,
				IsLead:          caster.Bool(true),
				IsFixedQuantity: caster.Bool(true),
				Duration:        caster.Int(14),
			},
			{
				RoleID:          6, // SM
				ModuleID:        m.ID,
				Quantity:        1,
				IsFixedQuantity: caster.Bool(true),
				Duration:        caster.Int(14),
			},
			{
				RoleID:          11, // UI
				ModuleID:        m.ID,
				Quantity:        1,
				IsFixedQuantity: caster.Bool(false),
				Duration:        caster.Int(12),
			},
			{
				RoleID:          12, // UX
				ModuleID:        m.ID,
				Quantity:        1,
				IsFixedQuantity: caster.Bool(false),
				Duration:        caster.Int(12),
			},
			{
				RoleID:          13, // BE
				ModuleID:        m.ID,
				Quantity:        2,
				IsFixedQuantity: caster.Bool(false),
				Duration:        caster.Int(12),
			},
			{
				RoleID:          14, // FE
				ModuleID:        m.ID,
				Quantity:        2,
				IsFixedQuantity: caster.Bool(false),
				Duration:        caster.Int(12),
			},
			{
				RoleID:          7, // ST
				ModuleID:        m.ID,
				Quantity:        1,
				IsFixedQuantity: caster.Bool(false),
				Duration:        caster.Int(12),
			},
		}, nil
	case ModuleDS:
		return []ModuleRole{
			{
				RoleID:          5, // PO
				ModuleID:        m.ID,
				Quantity:        1,
				IsLead:          caster.Bool(true),
				IsFixedQuantity: caster.Bool(true),
				Duration:        caster.Int(14),
			},
			{
				RoleID:          6, // SM
				ModuleID:        m.ID,
				Quantity:        1,
				IsFixedQuantity: caster.Bool(true),
				Duration:        caster.Int(14),
			},
			{
				RoleID:          11, // UI
				ModuleID:        m.ID,
				Quantity:        2,
				IsFixedQuantity: caster.Bool(false),
				Duration:        caster.Int(12),
			},
			{
				RoleID:          12, // UX
				ModuleID:        m.ID,
				Quantity:        2,
				IsFixedQuantity: caster.Bool(false),
				Duration:        caster.Int(12),
			},
			{
				RoleID:          13, // BE
				ModuleID:        m.ID,
				Quantity:        2,
				IsFixedQuantity: caster.Bool(false),
				Duration:        caster.Int(12),
			},
		}, nil
	}
}
